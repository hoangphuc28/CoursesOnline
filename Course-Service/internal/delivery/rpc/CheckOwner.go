package rpc

import (
	"context"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
)

func (hdl *coursesHandler) CheckOwner(ctx context.Context, req *Course.CheckOwnerRequest) (*Course.Response, error) {
	if err := hdl.uc.CheckOwner(req); err != nil {
		return &Course.Response{Error: HandleError(err)}, nil
	}
	return &Course.Response{}, nil
}
