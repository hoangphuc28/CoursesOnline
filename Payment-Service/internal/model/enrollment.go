package model

import (
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/pkg/common"
	"time"
)

type Enrollment struct {
	UserId         int       `gorm:"column:user_id"`
	CourseId       int       `gorm:"column:course_id"`
	EnrollmentDate time.Time `gorm:"column:enrollment_date;default:current_timestamp"`
}
type User struct {
	common.SQLModel
	FirstName string   `json:"firstName" gorm:"column:firstName"`
	LastName  string   `json:"lastName" gorm:"column:lastName"`
	Courses   []Course `gorm:"many2many:Enrollments;"`
}

func (Enrollment) TableName() string {
	return "enrollments"
}
