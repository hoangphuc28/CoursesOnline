package usecase

import (
	"fmt"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/internal/model"
)

func (uc *coursesUseCase) CreateLecture(req *Course.CreateLectureRequest) (*Course.CreateLectureResponse, error) {
	sectionId, err := uc.h.Decode(req.SectionId)
	if err != nil {
		return nil, err
	}
	res, err := uc.repo.NewLecture(model.Lecture{

		Title:     req.Title,
		SectionId: sectionId,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(res.Id)
	return &Course.CreateLectureResponse{
		LectureId: uc.h.Encode(res.Id),
	}, nil
}
func (uc *coursesUseCase) UpdateLecture(req *Course.UpdateLectureRequest) error {
	lectureId, err := uc.h.Decode(req.LectureId)
	if err != nil {
		return err
	}
	s, err := uc.repo.FindLectureWithCondition(map[string]any{"id": lectureId})
	if err != nil {
		return err
	}
	s.Title = req.Title
	s.Resource.Url = req.Resource.Url
	s.Resource.Duration = req.Resource.Duration
	if err = uc.repo.UpdateLecture(s, map[string]any{"id": lectureId}); err != nil {
		return err
	}
	return nil
}
func (uc *coursesUseCase) DeleteLecture(req *Course.DeleteLectureRequest) error {
	lectureId, err := uc.h.Decode(req.LectureId)
	if err != nil {
		return err
	}
	s, err := uc.repo.FindLectureWithCondition(map[string]any{"id": lectureId})
	if err != nil {
		return err
	}
	if err := uc.repo.DeleteLecture(s); err != nil {
		return err
	}
	return nil
}
