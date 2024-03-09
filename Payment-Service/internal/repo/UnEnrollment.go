package repo

import "github.com/hoangphuc28/CoursesOnline/Payment-Service/internal/model"

func (rp *paymentRepo) UnEnrollment(e *model.Enrollment) (*model.Enrollment, error) {
	if err := rp.db.Model(e).Where(e).Delete(&e).Error; err != nil {
		return nil, err
	}
	return e, nil
}
