package rpc

import (
	"context"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
)

func (hdl *coursesHandler) UpdateCourse(ctx context.Context, rq *Course.UpdateCourseRequest) (*Course.UpdateCourseResponse, error) {
	if err := hdl.uc.UpdateCourse(rq); err != nil {
		return &Course.UpdateCourseResponse{Error: HandleError(err)}, nil
	}

	return &Course.UpdateCourseResponse{}, nil
}
