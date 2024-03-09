package repo

import (
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/internal/model"
)

func (repo *paymentRepo) Refund(p *model.Payment) error {
	p.RefundStatus = "refunded"
	if err := repo.db.Save(&p).Error; err != nil {
		return err
	}
	return nil
}
