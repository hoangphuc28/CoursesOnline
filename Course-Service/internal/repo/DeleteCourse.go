package repo

import "github.com/hoangphuc28/CoursesOnline/Course-Service/internal/model"

func (r *coursesRepository) DeleteCourse(course model.Course) error {
	db := r.db.Begin()
	if err := db.Delete(&course).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
