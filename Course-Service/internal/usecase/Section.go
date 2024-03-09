package usecase

import (
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/internal/model"
)

func (uc *coursesUseCase) CreateSection(req *Course.CreateSectionRequest) (*Course.CreateSectionResponse, error) {
	courseId, err := uc.h.Decode(req.CourseId)
	if err != nil {
		return nil, err
	}
	res, err := uc.repo.NewSection(model.Section{
		Title:    req.Title,
		CourseId: courseId,
	})
	if err != nil {
		return nil, err
	}
	return &Course.CreateSectionResponse{
		SectionId: uc.h.Encode(res.Id),
	}, nil
}
func (uc *coursesUseCase) UpdateSection(req *Course.UpdateSectionRequest) error {
	sectionId, err := uc.h.Decode(req.SectionId)
	if err != nil {
		return err
	}
	s, err := uc.repo.FindSectionWithCondition(map[string]any{"id": sectionId})
	if err != nil {
		return err
	}
	s.Title = req.Title
	if err = uc.repo.UpdateSection(s, map[string]any{"id": sectionId}); err != nil {
		return err
	}
	return nil
}
func (uc *coursesUseCase) DeleteSection(req *Course.DeleteSectionRequest) error {
	sectionId, err := uc.h.Decode(req.SectionId)
	if err != nil {
		return err
	}
	s, err := uc.repo.FindSectionWithCondition(map[string]any{"id": sectionId})
	if err != nil {
		return err
	}
	if err := uc.repo.DeleteSection(s); err != nil {
		return err
	}
	return nil
}

func (uc *coursesUseCase) GetSections(request *Course.GetSectionsRequest) (*Course.GetSectionsResponse, error) {
	courseId, err := uc.h.Decode(request.CourseId)
	if err != nil {
		return nil, err
	}
	sections, err := uc.repo.FindSectionsWithCondition(map[string]any{"course_id": courseId})
	if err != nil {
		return nil, err
	}
	var res Course.GetSectionsResponse
	for _, s := range sections {
		var lectures []*Course.Lecture
		for _, l := range s.Lectures {
			lectures = append(lectures, &Course.Lecture{
				Id:    uc.h.Encode(l.Id),
				Title: l.Title,
				Resource: &Course.Resource{
					Url:      l.Resource.Url,
					Duration: l.Resource.Duration,
				},
				IsFree: l.IsFree,
			})
		}
		res.Sections = append(res.Sections, &Course.Section{
			Id:       uc.h.Encode(s.Id),
			Title:    s.Title,
			Lectures: lectures,
		})
	}
	return &res, nil
}
