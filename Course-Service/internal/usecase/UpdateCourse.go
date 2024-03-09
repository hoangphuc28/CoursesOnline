package usecase

import (
	"fmt"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/internal/model"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/pkg/common"
	"net/http"
)

func (uc *coursesUseCase) UpdateCourse(rq *Course.UpdateCourseRequest) error {
	courseIdDecoded, err := uc.h.Decode(rq.CourseId)
	if err != nil {
		return err
	}
	instructorIdDecoded, err := uc.h.Decode(rq.InstructorId)
	if err != nil {
		return err
	}
	course, err := uc.repo.GetCourse(courseIdDecoded)
	if err != nil {
		return err
	}
	if course.InstructorID != instructorIdDecoded {
		return common.NewCustomError(err, http.StatusNotFound, "Your are not an instructor of this course!")
	}
	var img common.Image
	if rq.Thumbnail != nil {
		img.Url = rq.Thumbnail.Url
		img.Width = rq.Thumbnail.Width
		img.Height = rq.Thumbnail.Height
	}
	idSubCate, err := uc.h.Decode(rq.SubCategory)
	if err != nil {
		fmt.Println(err)
		return err
	}
	course.Title = rq.Title
	course.Description = rq.Description
	course.Level = rq.Level
	course.Language = rq.Language
	course.PriceID = int(rq.Price)
	course.IsPublish = false
	course.Thumbnail = &img
	course.SubcategoryId = idSubCate
	course.Goals = rq.Goals
	course.Requirement = rq.Requirement
	var sections []model.Section
	for _, section := range rq.Sections {
		var lectures []model.Lecture
		for _, lecture := range section.Lectures {
			if lecture.Id != "" {
				idLecture, _ := uc.h.Decode(lecture.Id)
				err := uc.repo.UpdateLecture(&model.Lecture{
					SQLModel: common.SQLModel{
						Id: idLecture, // Set the desired ID here
					},
					Title:   lecture.Title,
					Content: lecture.Content,
					Status:  lecture.Status,
					Resource: model.Resource{
						Url:      lecture.Resource.Url,
						Duration: lecture.Resource.Duration,
					},
				}, map[string]any{"id": idLecture})
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("ok")
				}
			} else {
				lectures = append(lectures, model.Lecture{
					Title:   lecture.Title,
					Content: lecture.Content,
					Status:  lecture.Status,
					Resource: model.Resource{
						Url:      lecture.Resource.Url,
						Duration: lecture.Resource.Duration,
					},
				})
			}
		}
		var idSection int
		if section.Id != "" {
			idSection, _ = uc.h.Decode(section.Id)
			err := uc.repo.UpdateSection(&model.Section{
				SQLModel: common.SQLModel{Id: idSection},
				Title:    section.Title,
				Lectures: lectures,
			}, map[string]any{"id": idSection})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("ok")
			}
		} else {
			sections = append(sections, model.Section{
				Title:    section.Title,
				Lectures: lectures,
			})
		}

	}
	course.Sections = sections

	if err := uc.repo.UpdateCourse(&course); err != nil {
		return err
	}
	return nil
}
