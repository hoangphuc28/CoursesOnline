package rpc

import (
	"context"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Payment"
)

func (hdl paymentHandler) GetPayment(ctx context.Context, req *Payment.GetPaymentRequest) (*Payment.GetPaymentResponse, error) {
	res, err := hdl.uc.GetPayment(req)
	if err != nil {
		return &Payment.GetPaymentResponse{Error: HandleError(err)}, nil
	}
	return res, nil
}
