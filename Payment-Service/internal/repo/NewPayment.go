package repo

import "github.com/hoangphuc28/CoursesOnline/Payment-Service/internal/model"

func (rp *paymentRepo) NewPayment(payment *model.Payment) (*model.Payment, error) {
	db := rp.db.Begin()
	if err := db.Table(model.Payment{}.TableName()).Preload("PaymentCourses").Create(&payment).Error; err != nil {
		db.Rollback()
		return nil, err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return nil, err
	}
	return nil, nil
}
