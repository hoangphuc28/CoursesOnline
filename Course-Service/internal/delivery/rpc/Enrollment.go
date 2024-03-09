package https

import (
	"context"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
)

func (hdl *coursesHandler) Enrollment(ctx context.Context, request *Course.EnrollmentRequest) (*Course.EnrollmentResponse, error) {
	if err := hdl.uc.NewEnrollment(request.UserId, request.CourseId); err != nil {
		return &Course.EnrollmentResponse{Error: HandleError(err)}, nil
	}
	return &Course.EnrollmentResponse{}, nil
}
func (hdl *coursesHandler) GetEnrollments(ctx context.Context, rq *Course.GetEnrollmentsRequest) (*Course.GetEnrollmentsResponse, error) {
	res, err := hdl.uc.GetEnrollments(rq.UserId)
	if err != nil {
		return &Course.GetEnrollmentsResponse{
			Error: HandleError(err),
		}, nil
	}
	return res, nil
}
