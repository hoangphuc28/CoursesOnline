package rpc

import (
	"context"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Payment"
)

func (hdl *paymentHandler) Refund(ctx context.Context, rq *Payment.RefundRequest) (*Payment.RefundResponse, error) {
	if err := hdl.uc.Refund(rq); err != nil {
		return &Payment.RefundResponse{Error: HandleError(err)}, nil
	}

	return &Payment.RefundResponse{}, nil
}
