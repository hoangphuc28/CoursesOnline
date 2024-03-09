package https

import (
	"context"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
)

func (hdl *coursesHandler) CreateLecture(ctx context.Context, request *Course.CreateLectureRequest) (*Course.CreateLectureResponse, error) {
	res, err := hdl.uc.CreateLecture(request)
	if err != nil {
		return &Course.CreateLectureResponse{Error: HandleError(err)}, nil
	}
	return res, nil
}
func (hdl *coursesHandler) UpdateLecture(ctx context.Context, request *Course.UpdateLectureRequest) (*Course.Response, error) {
	if err := hdl.uc.UpdateLecture(request); err != nil {
		return &Course.Response{Error: HandleError(err)}, nil
	}
	return &Course.Response{}, nil
}
func (hdl *coursesHandler) DeleteLecture(ctx context.Context, request *Course.DeleteLectureRequest) (*Course.Response, error) {
	if err := hdl.uc.DeleteLecture(request); err != nil {
		return &Course.Response{Error: HandleError(err)}, nil
	}
	return &Course.Response{}, nil
}
