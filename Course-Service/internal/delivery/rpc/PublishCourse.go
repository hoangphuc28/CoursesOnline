package https

import (
	"context"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
)

func (hdl *coursesHandler) PublishCourse(ctx context.Context, req *Course.PublishCourseRequest) (*Course.PublishCourseResponse, error) {

	errorsValidate, err := hdl.uc.PublishCourse(ctx, req)
	if errorsValidate != nil {
		return &Course.PublishCourseResponse{ErrorsValidate: errorsValidate}, nil
	}

	if err != nil {
		return &Course.PublishCourseResponse{Error: HandleError(err)}, nil
	}
	return &Course.PublishCourseResponse{}, nil
}
