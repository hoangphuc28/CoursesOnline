package repo

import "github.com/hoangphuc28/CoursesOnline/Course-Service/internal/model"

func (c *coursesRepository) UpdateCourse(course *model.Course) error {
	var price model.Price
	if err := c.db.Table("Price").Where("id = ?", course.PriceID).First(&price).Error; err != nil {
		return err
	}
	course.Price = price
	db := c.db.Begin()
	if err := db.Save(course).Error; err != nil {
		db.Rollback()
		return err
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}

	return nil
}
