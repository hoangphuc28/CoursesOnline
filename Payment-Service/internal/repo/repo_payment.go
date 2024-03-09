package repo

import (
	"errors"
	"fmt"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/internal/model"
	"gorm.io/gorm"
)

type paymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) *paymentRepo {
	return &paymentRepo{db: db}
}
func (repo *paymentRepo) GetPayment(condition map[string]interface{}) (*model.Payment, error) {
	var res model.Payment

	if err := repo.db.Table(model.Payment{}.TableName()).Where(condition).Preload("PaymentCourses").First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}
func (repo *paymentRepo) FindData(condition map[string]interface{}) ([]model.Payment, error) {
	var data []model.Payment
	fmt.Println(condition)
	if err := repo.db.Table(model.Payment{}.TableName()).Where(condition).Preload("PaymentCourses").Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
func (repo *paymentRepo) GetInformationPaypal(instructorId int) (*model.Paypal, error) {
	var instructor model.Instructor

	if err := repo.db.Table(model.Instructor{}.TableName()).Where("id = ?", instructorId).First(&instructor).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	var pp *model.Paypal
	if err := repo.db.Table(model.Paypal{}.TableName()).Where("user_id = ?", instructor.UserId).First(&pp).Error; err != nil {
		return nil, err
	}
	return pp, nil
}

func (repo *paymentRepo) NewPaypalMethod(paypal model.Paypal) error {
	var pp model.Paypal
	db := repo.db.Begin()
	res := repo.db.Table(model.Paypal{}.TableName()).Where("user_id = ?", paypal.UserId).First(&pp)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		if err := db.Table(model.Paypal{}.TableName()).Delete(&pp).Error; err != nil {
			db.Rollback()
			return err
		}
	}
	if err := db.Table(model.Paypal{}.TableName()).Create(&paypal).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
