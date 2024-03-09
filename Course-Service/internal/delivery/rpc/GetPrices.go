package rpc

import (
	"context"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
)

func (hdl *coursesHandler) GetPrices(ctx context.Context, request *Course.GetPricesRequest) (*Course.GetPricesResponse, error) {
	res, err := hdl.uc.GetPrices()
	if err != nil {
		return &Course.GetPricesResponse{
			Error: HandleError(err),
		}, nil
	}
	return res, nil
}
