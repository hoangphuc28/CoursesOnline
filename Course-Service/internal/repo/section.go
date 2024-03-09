package repo

import "github.com/hoangphuc28/CoursesOnline/Course-Service/internal/model"

func (rp *coursesRepository) NewSection(section model.Section) (*model.Section, error) {
	if err := rp.db.Table(model.Section{}.TableName()).Create(&section).Error; err != nil {
		return nil, err
	}
	return &section, nil
}
func (rp *coursesRepository) UpdateSection(section *model.Section, condition map[string]any) error {
	if err := rp.db.Table(model.Section{}.TableName()).Where(condition).Updates(&section).Error; err != nil {
		return err
	}
	return nil
}
func (rp *coursesRepository) DeleteSection(s *model.Section) error {
	if err := rp.db.Table(model.Section{}.TableName()).Delete(&s).Error; err != nil {
		return err
	}
	return nil
}
func (rp *coursesRepository) FindSectionWithCondition(condition map[string]any) (*model.Section, error) {
	var s model.Section
	if err := rp.db.Table(model.Section{}.TableName()).Where(condition).First(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}
func (rp *coursesRepository) FindSectionsWithCondition(condition map[string]any) ([]model.Section, error) {
	var s []model.Section
	if err := rp.db.Table(model.Section{}.TableName()).Where(condition).Preload("Lectures.Resource").Find(&s).Error; err != nil {
		return nil, err
	}
	return s, nil
}
