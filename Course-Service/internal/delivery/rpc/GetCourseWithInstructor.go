package rpc

import (
	"context"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
)

func (hdl *coursesHandler) GetCourseWithInstructor(ctx context.Context, req *Course.GetCourseWithInstructorRequest) (*Course.GetCourseWithInstructorResponse, error) {

	res, err := hdl.uc.GetCoursesWithInstructor(req)
	if err != nil {

		return &Course.GetCourseWithInstructorResponse{Error: HandleError(err)}, nil
	}
	return res, nil
}
