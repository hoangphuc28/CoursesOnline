package usecase

import (
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/pkg/common"
	"net/http"
)

func (uc *coursesUseCase) DeleteCourse(req *Course.DeleteCourseRequest) error {
	userIdDecoded, err := uc.h.Decode(req.UserId)
	if err != nil {
		return err
	}
	courseIdDecoded, err := uc.h.Decode(req.CourseId)
	if err != nil {
		return err
	}
	course, err := uc.repo.GetCourse(courseIdDecoded)
	if err != nil {
		return err
	}
	if course.InstructorID != userIdDecoded {
		return common.NewCustomError(err, http.StatusNotFound, "Your are not allow!")
	}
	if err := uc.repo.DeleteCourse(course); err != nil {
		return common.NewCustomError(err, http.StatusBadRequest, "Unable delete course!")
	}
	return nil
}
