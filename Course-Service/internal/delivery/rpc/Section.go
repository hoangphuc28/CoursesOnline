package rpc

import (
	"context"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
)

func (hdl *coursesHandler) GetSections(ctx context.Context, req *Course.GetSectionsRequest) (*Course.GetSectionsResponse, error) {
	res, err := hdl.uc.GetSections(req)
	if err != nil {
		return &Course.GetSectionsResponse{Error: HandleError(err)}, nil
	}
	return res, nil
}
func (hdl *coursesHandler) CreateSection(ctx context.Context, request *Course.CreateSectionRequest) (*Course.CreateSectionResponse, error) {
	res, err := hdl.uc.CreateSection(request)
	if err != nil {
		return &Course.CreateSectionResponse{Error: HandleError(err)}, nil
	}
	return res, nil
}
func (hdl *coursesHandler) UpdateSection(ctx context.Context, request *Course.UpdateSectionRequest) (*Course.Response, error) {
	if err := hdl.uc.UpdateSection(request); err != nil {
		return &Course.Response{Error: HandleError(err)}, nil
	}
	return &Course.Response{}, nil
}
func (hdl *coursesHandler) DeleteSection(ctx context.Context, request *Course.DeleteSectionRequest) (*Course.Response, error) {
	if err := hdl.uc.DeleteSection(request); err != nil {
		return &Course.Response{Error: HandleError(err)}, nil
	}
	return &Course.Response{}, nil
}
