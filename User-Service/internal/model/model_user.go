package model

import (
	"github.com/Zhoangp/User-Service/pkg/common"
	"github.com/Zhoangp/User-Service/pkg/utils"
	"gorm.io/gorm"
	"time"
)

type Users struct {
	common.SQLModel
	Email        string        `json:"email" gorm:"column:email"`
	Password     string        `json:"password" gorm:"column:password"`
	FirstName    string        `json:"firstName" gorm:"column:firstName"`
	LastName     string        `json:"lastName" gorm:"column:lastName"`
	Phone        string        `json:"phone" gorm:"column:phoneNumber"`
	Role         string        `json:"role" gorm:"column:role"`
	Address      string        `json:"address" gorm:"column:address"`
	IsInstructor bool          `json:"is_instructor"`
	Avatar       *common.Image `gorm:"column:picture"`
}
type UserRegister struct {
	Credential   string `json:"credential"`
	FirstName    string `json:"firstName" gorm:"column:firstName"`
	LastName     string `json:"lastName" gorm:"column:lastName"`
	Phone        string `json:"phoneNumber" gorm:"column:phoneNumber"`
	Role         string `json:"role" gorm:"column:role"`
	Address      string `json:"address"`
	IsInstructor int    `json:"is_instructor"`
}
type UserLogin struct {
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}
type UserChangePassword struct {
	Email       string `json:"email" gorm:"column:email"`
	OldPassword string `json:"password" gorm:"column:password"`
	NewPass     string `json:"newpassword"`
}
type Instructor struct {
	UserId       int            `gorm:"column:id"`
	FakeId       string         `json:"id" gorm:"-"`
	CreatedAt    time.Time      `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt    time.Time      `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
	Website      string         `gorm:"column:website" valid:"required~Website không được để trống"`
	LinkedIn     string         `gorm:"column:linkedin" valid:"required~LinkedIn không được để trống"`
	Youtube      string         `gorm:"column:youtube" valid:"required~Youtube không được để trống"`
	Bio          string         `gorm:"column:bio" valid:"required"`
	Paypal       Paypal         `gorm:"foreignKey:UserId;references:id" valid:"required~Tài khoản paypal không được để trống"`
	NumStudents  string         `gorm:"column:num_students"`
	NumReviews   string         `gorm:"column:num_reviews"`
	TotalCourses string         `gorm:"column:total_courses"`
	User         Users          `gorm:"foreignKey:id;references:UserId"`
}

type Paypal struct {
	common.SQLModel
	UserId   int    `json:"userId" gorm:"column:user_id"`
	Email    string `json:"email" gorm:"column:email"`
	PaypalId string `json:"payerId" gorm:"column:paypal_id"`
}

func (Instructor) TableName() string {
	return "Instructor"
}
func (Users) TableName() string {
	return "Users"
}

func (u *Users) GetUserId() int {
	return u.Id
}

func (u *Users) GetUserEmail() string {
	return u.Email
}

func (u *Users) GetUserRole() string {
	return u.Role
}
func (u *Users) PrepareCreate() error {
	passHashed, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = passHashed
	u.Role = "guest"
	return nil
}
