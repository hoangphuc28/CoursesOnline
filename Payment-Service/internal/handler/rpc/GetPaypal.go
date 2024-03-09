package rpc

import (
	"context"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Payment"
)

func (hdl *paymentHandler) GetPaypal(ctx context.Context, req *Payment.GetPayalRequest) (*Payment.GetPayalResponse, error) {
	res, err := hdl.uc.GetPaypal(req)
	if err != nil {
		return &Payment.GetPayalResponse{Error: HandleError(err)}, nil
	}
	return res, nil
}
