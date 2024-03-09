package model

import "github.com/hoangphuc28/CoursesOnline/Cart-Service/pkg/common"

type Cart struct {
	common.SQLModel
	UserId  int      `gorm:"column:user_id"`
	Courses []Course `gorm:"many2many:Cart_Courses;"`
}
type CartCourse struct {
	CartId   int `gorm:"column:cart_id"`
	CourseId int `gorm:"column:course_id"`
}
type Course struct {
	common.SQLModel
	Title             string        `json:"title" gorm:"column:title"`
	CourseDescription string        `json:"description" gorm:"column:description"`
	CourseLevel       string        `json:"level"  gorm:"column:level"`
	CourseLanguage    string        `json:"language" gorm:"column:language"`
	PriceID           int           `valid:"required~Price is required" gorm:"column:price_id"`
	Price             Price         `gorm:"foreignKey:PriceID;references:id"`
	CourseCurrency    string        `json:"currency" gorm:"column:currency"`
	CourseDiscount    float32       `json:"discount" gorm:"column:discount"`
	CourseDuration    string        `json:"duration" gorm:"column:duration"`
	CourseStatus      string        `json:"status" gorm:"column:status"`
	CourseRating      float32       `json:"rating" gorm:"column:rating"`
	InstructorID      int           `json:"instructor_id" gorm:"column:instructor_id"`
	CourseThumbnail   *common.Image `json:"thumbnail" gorm:"column:thumbnail"`
	SubCategoryId     int           `gorm:"column:subCategory_id"`
}
type Price struct {
	common.SQLModel
	Value    string `json:"value" gorm:"column:value"`
	Currency string `json:"currency" gorm:"column:currency"`
}

func (CartCourse) TableName() string {
	return "cart_courses"
}
func (Course) TableName() string {
	return "Course"
}

func (Cart) TableName() string {
	return "Cart"
}
func (Price) TableName() string {
	return "Price"
}
