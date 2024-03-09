package https

import (
	"context"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
)

func (hdl *coursesHandler) GetCourseContent(ctx context.Context, request *Course.GetCourseContentRequest) (*Course.GetCourseContentResponse, error) {
	res, err := hdl.uc.GetCourseContent(request.UserId, request.CourseId)
	if err != nil {
		return &Course.GetCourseContentResponse{
			Error: HandleError(err),
		}, nil
	}
	return res, nil
}
