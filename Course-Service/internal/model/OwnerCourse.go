package model

type OwnerCourse struct {
	UserId   int `gorm:"column:user_id"`
	CourseId int `gorm:"column:course_id"`
}

func (OwnerCourse) TableName() string {
	return "owner_courses"
}
