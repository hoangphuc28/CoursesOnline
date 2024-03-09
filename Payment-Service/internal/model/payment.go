package model

import "github.com/hoangphuc28/CoursesOnline/Payment-Service/pkg/common"

type Paypal struct {
	common.SQLModel
	UserId   int    `json:"userId" gorm:"column:user_id"`
	Email    string `json:"email" gorm:"column:email"`
	PaypalId string `json:"payerId" gorm:"column:paypal_id"`
}
type Payment struct {
	common.SQLModel
	UserId         int             `json:"userId" gorm:"column:user_id"`
	Total          string          `gorm:"column:total"`
	RefundStatus   string          `gorm:"column:refund_status"`
	RefundUrl      string          `gorm:"column:refund_url"`
	PaymentCourses []PaymentCourse `gorm:"foreignKey:PaymentId;references:Id;constraint:OnDelete:CASCADE"`
}
type PaymentCourse struct {
	PaymentId int     `gorm:"column:payment_id"`
	CourseId  int     `gorm:"column:course_id"`
	Price     string  `gorm:"column:price"`
	Discount  float64 `gorm:"column:discount"`
	Amount    string  `gorm:"column:amount"`
	Currency  string  `gorm:"column:currency"`
	Title     string  `gorm:"column:title"`
}
type Course struct {
	common.SQLModel
	Title          string  `json:"title" gorm:"column:title"`
	Price          string  `gorm:"foreignKey:PriceID;references:id"`
	CourseCurrency string  `json:"currency" gorm:"column:currency"`
	CourseDiscount float32 `json:"discount" gorm:"column:discount"`
}
type Price struct {
	common.SQLModel
	Value    string `json:"value" gorm:"column:value"`
	Currency string `json:"currency" gorm:"column:currency"`
}

func (Paypal) TableName() string {
	return "Paypal"
}
func (Payment) TableName() string {
	return "Payment"
}
func (PaymentCourse) TableName() string {
	return "Payment_Course"
}

type Instructor struct {
	common.SQLModel
	UserId int `gorm:"column:user_id"`
}

func (Instructor) TableName() string {
	return "Instructor"
}
