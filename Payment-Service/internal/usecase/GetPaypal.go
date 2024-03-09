package usecase

import (
	"errors"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Payment"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/pkg/common"
)

func (uc *paymentUseCase) GetPaypal(req *Payment.GetPayalRequest) (*Payment.GetPayalResponse, error) {
	userIdDecoded, err := uc.h.Decode(req.UserId)
	if err != nil {
		return nil, err
	}
	paypal, err := uc.repo.GetInforPaypal(userIdDecoded)
	if paypal == nil {
		return nil, common.NewCustomError(errors.New("Paypal account not found"), "Paypal account not found")
	}
	return &Payment.GetPayalResponse{
		Email: paypal.Email,
	}, nil
}
