package repo

import (
	"fmt"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/internal/model"
)

func (repo *paymentRepo) GetInforPaypal(userId int) (*model.Paypal, error) {
	var paypal model.Paypal
	if err := repo.db.Table(model.Paypal{}.TableName()).Where("user_id = ?", userId).First(&paypal).Error; err != nil {
		return nil, err
	}
	fmt.Println(paypal)
	return &paypal, nil
}
