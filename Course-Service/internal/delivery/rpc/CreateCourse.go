package https

import (
	"context"
	"fmt"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
)

func (hdl *coursesHandler) CreateCourse(ctx context.Context, rq *Course.CreateCourseRequest) (*Course.CreateCourseResponse, error) {
	fmt.Println(rq)

	courseId, err := hdl.uc.NewCourse(rq, rq.InstructorId)

	if err != nil {
		return &Course.CreateCourseResponse{
			Error: HandleError(err),
		}, nil
	}
	return &Course.CreateCourseResponse{
		CourseId: courseId,
	}, nil
}
