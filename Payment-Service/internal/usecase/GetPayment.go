package usecase

import (
	"fmt"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Payment"
	"math"
	"strconv"
)

func (uc *paymentUseCase) GetPayment(req *Payment.GetPaymentRequest) (*Payment.GetPaymentResponse, error) {

	userIdDecoded, err := uc.h.Decode(req.UserId)
	if err != nil {
		return nil, err
	}
	payments, err := uc.repo.FindData(map[string]any{"user_id": userIdDecoded})
	if err != nil {
		return nil, err
	}
	var res Payment.GetPaymentResponse
	for _, payment := range payments {
		pc, _ := strconv.ParseFloat(payment.Total, 64)

		var item Payment.Payment
		item.PaymentId = uc.h.Encode(payment.Id)
		item.Total = fmt.Sprintf("%f", math.Ceil(pc*23252.2009))
		item.PaymentDate = payment.CreatedAt.Format("2006-01-02 15:04:05")
		fmt.Println(payment.RefundStatus)
		ok, err := CompareTime(payment.CreatedAt)
		if err != nil {
			return nil, err
		}
		if !ok {
			item.RefundStatus = "expiredNoRefund"
		} else {
			item.RefundStatus = payment.RefundStatus
		}
		fmt.Println(item.RefundStatus)
		for _, paymentCourse := range payment.PaymentCourses {
			item.ListItem = append(item.ListItem, &Payment.ListItemPayment{
				Title:    paymentCourse.Title,
				Price:    paymentCourse.Price,
				Discount: fmt.Sprintf("%f", paymentCourse.Discount),
				Amount:   paymentCourse.Amount,
			})
		}
		res.Payment = append(res.Payment, &item)
	}

	return &res, nil
}
