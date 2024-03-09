package model

import (
	"github.com/hoangphuc28/CoursesOnline/Course-Service/pkg/common"
	"time"
)

type Course struct {
	common.SQLModel
	Title       string `valid:"required~Course title is required" gorm:"column:title"`
	Description string `valid:"required~Course description is required" gorm:"column:description"`
	Level       string `valid:"required~Level is required" gorm:"column:level"`
	Language    string `valid:"required~Language is required" gorm:"column:language"`

	PriceID int   `valid:"required~Price is required" gorm:"column:price_id"`
	Price   Price `gorm:"foreignKey:PriceID;references:id"`

	IsPublish     bool       `gorm:"column:is_publish"`
	PublishedTime *time.Time `gorm:"column:published_time"`

	InstructorID int           `gorm:"column:instructor_id"`
	Thumbnail    *common.Image `valid:"required~Thumbnail is required" gorm:"column:thumbnail"`
	Sections     []Section     `valid:"required~Sections is required" gorm:"foreignKey:CourseId;references:Id"`

	Goals       string `valid:"required~Goals is required"`
	Requirement string `valid:"required~Requirement is required" gorm:"column:requirement"`

	NumReviews     int       `gorm:"column:num_reviews"`
	NumSubscribers int       `gorm:"column:num_subscribers"`
	Rating         float32   `gorm:"column:rating"`
	TotalLectures  int       `gorm:"column:total_lectures"`
	TotalSections  int       `gorm:"column:total_sections"`
	TotalLength    time.Time `gorm:"-"`

	SubcategoryId int `valid:"required~Course subcategory is required" gorm:"column:SubCategory_id"`
}
type Coupon struct {
	common.SQLModel
	Code       string    `gorm:"column:code"`
	Percentage float32   `gorm:"column:percentage"`
	StartDate  time.Time `gorm:"column:start_date"`
	ExpiryDate time.Time `gorm:"column:expiry_date"`
}
type Price struct {
	common.SQLModel
	Value    string `json:"value" gorm:"column:value"`
	Currency string `json:"currency" gorm:"column:currency"`
}
type Instructor struct {
	common.SQLModel
	TotalStudent int    `gorm:"column:totalStudent"`
	TotalCourse  int    `gorm:"column:totalCourse"`
	TotalReview  int    `gorm:"column:totalReview"`
	About        string `gorm:"column:about"`
}

type Section struct {
	common.SQLModel
	Title            string    `valid:"required~Section title is required" gorm:"column:title"`
	CourseId         int       `gorm:"column:course_id"`
	NumberOfLectures int       `gorm:"-"`
	Lectures         []Lecture `json:"lectures" gorm:"foreignKey:SectionId;references:Id;constraint:OnDelete:CASCADE;"`
}
type Lecture struct {
	common.SQLModel
	Title     string   `valid:"required~Lecture title is required"  gorm:"column:title"`
	Content   string   `valid:"required~Lecture content is required"  gorm:"content"`
	Status    string   `gorm:"-"`
	SectionId int      `json:"section_id" gorm:"column:section_id"`
	IsFree    bool     `gorm:"column:is_free"`
	Resource  Resource `valid:"required~Lecture resource is required" gorm:"foreignKey:LectureId;references:id;constraint:OnDelete:CASCADE;"`
}
type Resource struct {
	common.SQLModel
	Url       string `valid:"required~Url resource is required" gorm:"column:url"`
	Duration  string `gorm:"column:duration"`
	LectureId int    `gorm:"lecture_id"`
}

func (Resource) TableName() string {
	return "Resource"
}
func (Lecture) TableName() string {
	return "Lectures"
}
func (Section) TableName() string {
	return "Section"
}
func (Course) TableName() string {
	return "Course"
}
func (Price) TableName() string {
	return "Price"
}
func (c *Course) PrepareCreate() {
	if c.Level == "" {
		c.Level = "All Level"
	}
	if c.Language == "" {
		c.Language = "en"
	}

}
