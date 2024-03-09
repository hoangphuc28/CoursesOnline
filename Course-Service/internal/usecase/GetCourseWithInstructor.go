package usecase

import "github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"

func (uc *coursesUseCase) GetCoursesWithInstructor(rq *Course.GetCourseWithInstructorRequest) (*Course.GetCourseWithInstructorResponse, error) {
	userIdDecoded, err := uc.h.Decode(rq.UserId)
	if err != nil {
		return nil, err
	}
	courses, err := uc.repo.FindDataWithCondition(map[string]any{"instructor_id": userIdDecoded})
	if err != nil {
		return nil, err
	}
	var res Course.GetCourseWithInstructorResponse
	for _, course := range courses {
		var img Course.Image
		if course.Thumbnail != nil {
			img = Course.Image{
				Url:    course.Thumbnail.Url,
				Width:  course.Thumbnail.Width,
				Height: course.Thumbnail.Height,
			}
		}
		res.Courses = append(res.Courses, &Course.Course{
			Id:        uc.h.Encode(course.Id),
			Title:     course.Title,
			Thumbnail: &img,
		})
	}
	return &res, nil
}
