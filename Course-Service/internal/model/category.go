package model

import "github.com/hoangphuc28/CoursesOnline/Course-Service/pkg/common"

type Category struct {
	common.SQLModel
	Name          string        `gorm:"name"`
	SubCategories []SubCategory `gorm:"foreignKey:CategoryId;"`
}
type SubCategory struct {
	common.SQLModel
	CategoryId int
	Name       string `gorm:"name"`
}

func (Category) TableName() string {
	return "Category"
}
func (SubCategory) TableName() string {
	return "Sub_Category"
}
