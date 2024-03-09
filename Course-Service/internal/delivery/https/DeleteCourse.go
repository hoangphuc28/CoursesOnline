package https

import (
	"context"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
)

func (hdl *coursesHandler) DeleteCourse(ctx context.Context, req *Course.DeleteCourseRequest) (*Course.DeleteCourseResponse, error) {
	if err := hdl.uc.DeleteCourse(req); err != nil {
		return &Course.DeleteCourseResponse{Error: HandleError(err)}, nil
	}
	return &Course.DeleteCourseResponse{}, nil
}
