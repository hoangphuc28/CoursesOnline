package repo

import "github.com/hoangphuc28/CoursesOnline/Course-Service/internal/model"

func (r coursesRepository) PublishCourse(course *model.Course) error {
	if err := r.db.Model(course).Update("is_publish", true).Error; err != nil {
		return err
	}
	return nil
}
