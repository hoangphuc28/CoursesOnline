package repo

import (
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/internal/model"
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/pkg/common"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) NewUsers(data *model.Users) error {
	db := repo.db.Begin()
	if err := db.Table(model.Users{}.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil

}
func (repo *UserRepository) FindDataWithCondition(conditions map[string]any) (*model.Users, error) {
	var user model.Users
	if err := repo.db.Table(model.Users{}.TableName()).Where(conditions).First(&user).Error; err != nil {
		return nil, common.ErrEntityNotFound("User-Service", err)
	}
	return &user, nil
}
func (repo *UserRepository) UpdateUser(user *model.Users, newInformation map[string]any) error {
	if err := repo.db.Model(&user).Updates(newInformation).Error; err != nil {
		return err
	}
	return nil
}
