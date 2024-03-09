package repo

import (
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/internal/model"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/pkg/common"
	"gorm.io/gorm"
	"net/http"
)

type coursesRepository struct {
	db *gorm.DB
}

func NewCoursesRepository(db *gorm.DB) *coursesRepository {
	return &coursesRepository{db}
}
func (c *coursesRepository) Create(course *model.Course) error {
	db := c.db.Begin()
	if err := db.Table(model.Course{}.TableName()).Create(course).Error; err != nil {
		fmt.Println(err)
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
func (c *coursesRepository) GetCourses(limit int, page int) (res []model.Course, total int64, err error) {
	db := c.db.Model(&res).Where("is_publish = ?", true)
	db.Count(&total)
	if limit <= 0 {
		db.Limit((int)(total)).Preload("Price").Find(&res)
	} else {
		db.Limit(limit).Offset(limit * (page - 1)).Preload("Price").Find(&res)
	}

	err = nil
	return
}
func (c *coursesRepository) GetCourse(id int) (res model.Course, err error) {
	db := c.db.Model(&res)

	if err = db.Where("id = ?", id).Preload("Price").Preload("Sections.Lectures.Resource").First(&res).Error; err != nil {
		return
	}
	return
}

func (c *coursesRepository) GetAllCategories() ([]model.Category, error) {
	var categories []model.Category
	if err := c.db.Table(model.Category{}.TableName()).Preload("SubCategories").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
func (c *coursesRepository) NewCourse(course *model.Course) (*int, error) {
	db := c.db.Begin()
	if err := db.Table(model.Course{}.TableName()).Create(&course).Error; err != nil {
		db.Rollback()
		return nil, err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return nil, err
	}
	return &course.Id, nil
}
func (c *coursesRepository) GetPrices() (res []model.Price, err error) {
	if err = c.db.Table(model.Price{}.TableName()).Find(&res).Error; err != nil {
		return nil, err
	}
	return
}
func (c *coursesRepository) GetCourseContent(e *model.Enrollment) (*model.Course, error) {
	if err := c.db.Table(model.Enrollment{}.TableName()).Where("user_id = ? AND course_id = ?", e.UserId, e.CourseId).First(&model.Enrollment{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.NewCustomError(err, http.StatusNotFound, "Users has not enrolled in this course yet!")
		}
		return nil, err
	}
	course, err := c.GetCourse(e.CourseId)
	if err != nil {
		return nil, err
	}
	return &course, nil
}
func (c *coursesRepository) NewEnrollment(enrollment *model.Enrollment) error {
	//course, err := c.GetCourse(enrollment.CourseId)
	//if err != nil {
	//	return err
	//}
	//if course.IsPaid {
	//	return common.NewCustomError(err, http.StatusConflict, "This course is not free!")
	//}
	db := c.db.Begin()
	if err := db.Table(model.Enrollment{}.TableName()).Create(&enrollment).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return common.NewCustomError(err, http.StatusConflict, "Users enrolled in this course!")
		}
		db.Rollback()
		return err
	}
	db.Commit()

	return nil
}
func (c *coursesRepository) GetEnrollments(u *model.User) ([]model.Course, error) {
	var courses []model.Course
	var err error
	if err = c.db.Model(u).Association("Courses").Find(&courses); err != nil {
		return nil, err
	}
	return courses, nil
}
