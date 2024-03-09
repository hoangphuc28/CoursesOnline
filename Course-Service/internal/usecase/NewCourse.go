package usecase

import (
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/internal/model"
)

func (uc *coursesUseCase) NewCourse(rq *Course.CreateCourseRequest, instructorId string) (string, error) {
	instructorIdDecoded, err := uc.h.Decode(instructorId)
	if err != nil {
		return "", err
	}
	SubCategoryIdDecoded, err := uc.h.Decode(rq.SubCategoryId)
	if err != nil {
		return "", err
	}
	course := model.Course{
		Title:         rq.Title,
		SubcategoryId: SubCategoryIdDecoded,
		InstructorID:  instructorIdDecoded,
		PriceID:       1,
	}
	course.PrepareCreate()

	courseId, err := uc.repo.NewCourse(&course)
	if err != nil {
		return "", err
	}
	courseIdEncoded := uc.h.Encode(*courseId)
	if err != nil {
		return "", err
	}
	return courseIdEncoded, nil
}
