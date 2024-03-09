package usecase

import (
	"errors"
	"fmt"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Payment"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/internal/model"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/pkg/common"
	"time"
)

func CompareTime(timePayment time.Time) (bool, error) {
	formattedTime := timePayment.Format("2006-01-02 15:04:05")

	layout := "2006-01-02 15:04:05"
	t1, err := time.Parse(layout, formattedTime)
	if err != nil {
		fmt.Println("Lỗi chuyển đổi thời gian:", err)
		return false, err
	}
	currentTime := time.Now()
	currentTimeFormated := currentTime.Format("2006-01-02 15:04:05")
	t2, err := time.Parse(layout, currentTimeFormated)
	if err != nil {
		fmt.Println("Lỗi chuyển đổi thời gian:", err)
		return false, err
	}
	duration := t2.Sub(t1)
	if duration < 0 {
		duration = -duration
	}
	if duration >= 30*time.Minute {
		return false, nil
	}
	return true, nil
}
func (uc *paymentUseCase) Refund(rq *Payment.RefundRequest) error {
	paymentId, err := uc.h.Decode(rq.PaymentId)
	if err != nil {
		return err
	}
	userId, err := uc.h.Decode(rq.UserId)
	if err != nil {
		return err
	}

	payment, err := uc.repo.GetPayment(map[string]any{"id": paymentId})
	if err != nil {
		return err
	}
	fmt.Println(payment)

	ok, err := CompareTime(payment.CreatedAt)
	if err != nil {
		return err
	}
	if !ok {
		return common.NewCustomError(errors.New("expired refund"), "expired refund")
	}
	if err := uc.repo.Refund(payment); err != nil {
		return err
	}

	for _, i := range payment.PaymentCourses {
		_, err := uc.repo.UnEnrollment(&model.Enrollment{
			UserId:   userId,
			CourseId: i.CourseId,
		})
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	uc.paypal.Refund(payment.RefundUrl)
	return nil
}
