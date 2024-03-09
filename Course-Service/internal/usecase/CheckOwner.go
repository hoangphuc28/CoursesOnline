package usecase

import (
	"errors"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/pkg/common"
)

func (uc *coursesUseCase) CheckOwner(rq *Course.CheckOwnerRequest) error {
	courseId, err := uc.h.Decode(rq.CourseId)
	if err != nil {
		return err
	}
	userId, err := uc.h.Decode(rq.UserId)
	if err != nil {
		return err
	}
	course, err := uc.repo.GetCourse(courseId)
	if err != nil {
		return err
	}
	if course.InstructorID != userId {
		return common.NewCustomError(errors.New("denied permission"), 403, "denied permission")
	}
	return nil
}
