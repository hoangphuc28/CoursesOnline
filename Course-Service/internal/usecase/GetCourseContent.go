package usecase

import (
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/User"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/internal/model"
	"strconv"
)

func (uc *coursesUseCase) GetCourseContent(userId, courseId string) (*Course.GetCourseContentResponse, error) {
	userIdDecoded, err := uc.h.Decode(userId)
	if err != nil {
		return nil, err
	}
	courseIdDecoded, err := uc.h.Decode(courseId)
	if err != nil {
		return nil, err
	}
	course, err := uc.repo.GetCourseContent(&model.Enrollment{
		UserId:   userIdDecoded,
		CourseId: courseIdDecoded,
	})
	if err != nil {
		return nil, err
	}
	course.FakeId = uc.h.Encode(course.Id)
	var sections []*Course.Section
	for _, i := range course.Sections {
		var lectures []*Course.Lecture
		for _, j := range i.Lectures {
			j.FakeId = uc.h.Encode(j.Id)
			lectures = append(lectures, &Course.Lecture{
				Id:      j.FakeId,
				Title:   j.Title,
				Content: j.Content,
				Resource: &Course.Resource{
					Url:      j.Resource.Url,
					Duration: j.Resource.Duration,
				},
			})
		}
		i.FakeId = uc.h.Encode(i.Id)
		sections = append(sections, &Course.Section{
			Id:               i.FakeId,
			Title:            i.Title,
			NumberOfLectures: int32(i.NumberOfLectures),
			Lectures:         lectures,
		})
	}

	res := &Course.Course{
		Id:    course.FakeId,
		Title: course.Title,
		Instructor: &User.Instructor{
			Id: uc.h.Encode(course.InstructorID),
		},
		Sections:   sections,
		NumReviews: strconv.Itoa(course.NumReviews),
	}

	return &Course.GetCourseContentResponse{
		Course: res,
	}, nil
}
