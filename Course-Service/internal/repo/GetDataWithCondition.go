package repo

import (
	"github.com/hoangphuc28/CoursesOnline/Course-Service/internal/model"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/pkg/common"
	"net/http"
)

func (repo *coursesRepository) FindDataWithCondition(conditions map[string]any) ([]model.Course, error) {
	var courses []model.Course
	if err := repo.db.Table(model.Course{}.TableName()).Where(conditions).Find(&courses).Error; err != nil {
		return nil, common.NewCustomError(err, http.StatusNotFound, "Does not have any course!")
	}
	return courses, nil
}
