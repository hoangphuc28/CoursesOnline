package repo

import "github.com/hoangphuc28/CoursesOnline/Course-Service/internal/model"

func (rp *coursesRepository) NewLecture(lecture model.Lecture) (*model.Lecture, error) {
	if err := rp.db.Table(model.Lecture{}.TableName()).Create(&lecture).Error; err != nil {
		return nil, err
	}
	return &lecture, nil
}
func (rp *coursesRepository) UpdateLecture(lecture *model.Lecture, condition map[string]any) error {
	if err := rp.db.Table(model.Lecture{}.TableName()).Where(condition).Updates(&lecture).Error; err != nil {
		return err
	}
	return nil
}
func (rp *coursesRepository) DeleteLecture(s *model.Lecture) error {
	if err := rp.db.Table(model.Lecture{}.TableName()).Delete(&s).Error; err != nil {
		return err
	}
	return nil
}
func (rp *coursesRepository) FindLectureWithCondition(condition map[string]any) (*model.Lecture, error) {
	var s model.Lecture
	if err := rp.db.Table(model.Lecture{}.TableName()).Where(condition).First(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}
