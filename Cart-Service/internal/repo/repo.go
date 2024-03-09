package repo

import (
	"github.com/hoangphuc28/CoursesOnline/Cart-Service/internal/model"
	"gorm.io/gorm"
)

type cartRepo struct {
	db *gorm.DB
}

func NewCartRepo(db *gorm.DB) *cartRepo {
	return &cartRepo{db: db}
}
func (r *cartRepo) NewCart(userId int) (*model.Cart, error) {
	cart := &model.Cart{
		UserId: userId,
	}
	if err := r.db.Table(model.Cart{}.TableName()).Create(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}
func (rp cartRepo) FindDataWithCondition(condition map[string]any) (*model.Cart, error) {
	var cart *model.Cart
	err := rp.db.Table(model.Cart{}.TableName()).Where(condition).Preload("Courses.Price").First(&cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}
func (rp cartRepo) FindCartWithUser(userId int) (int, error) {
	var res model.Cart
	if err := rp.db.Table(model.Cart{}.TableName()).Where("user_id = ?", userId).First(&res).Error; err != nil {
		return -1, err
	}
	return res.Id, nil
}

func (rp cartRepo) AddItemToCart(item *model.CartCourse) error {
	db := rp.db.Begin()
	if err := db.Table(model.CartCourse{}.TableName()).Create(item).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil

}

//	func (rp cartRepo) CreateNewCart(userId int) (*model.Cart, error) {
//		db := rp.db.Begin()
//		cart := model.Cart{
//			UserId: userId,
//		}
//		if err := db.Table(model.Cart{}.TableName()).Create(&cart).Error; err != nil {
//			db.Rollback()
//			return nil, err
//		}
//		if err := db.Commit().Error; err != nil {
//			db.Rollback()
//			return nil, err
//		}
//		return &cart, nil
//	}
func (rp cartRepo) ResetCart(cartId int) error {
	db := rp.db.Begin()
	var cart model.Cart
	if err := db.Table(model.Cart{}.TableName()).Where("id = ?", cartId).Preload("Courses").First(&cart).Error; err != nil {
		return err
	}
	userId := cart.UserId
	if _, err := rp.NewCart(userId); err != nil {
		db.Rollback()
		return err
	}
	if err := db.Table(model.Cart{}.TableName()).Delete(&cart).Error; err != nil {
		db.Rollback()
		return nil
	}
	if err := db.Commit(); err != nil {
		db.Rollback()
	}
	return nil
}
func (rp cartRepo) RemoveItemToCart(item *model.CartCourse) error {
	db := rp.db.Begin()
	if err := db.Table(model.CartCourse{}.TableName()).Where("cart_id = ? and course_id = ?", item.CartId, item.CourseId).Delete(item).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil

}
