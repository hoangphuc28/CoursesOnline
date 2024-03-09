package repo

import (
	"github.com/hoangphuc28/CoursesOnline/User-Service/internal/model"
	"github.com/hoangphuc28/CoursesOnline/User-Service/pkg/common"
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
func (repo *UserRepository) GetInstructor(condition map[string]any) (*model.Instructor, error) {
	var instructor model.Instructor
	if err := repo.db.Table(model.Instructor{}.TableName()).Where(condition).Preload("User").First(&instructor).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.NewCustomError(err, "User is not an instructor!")

		}
		return nil, err
	}
	return &instructor, nil
}
func (repo *UserRepository) NewInstructor(user *model.Users, intructor *model.Instructor) error {
	db := repo.db.Begin()
	if err := db.Model(&user).Update("role", "instructor").Error; err != nil {
		db.Rollback()
		return err
	}

	if err := db.Table(model.Instructor{}.TableName()).Create(intructor).Error; err != nil {
		return err
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}

	return nil

}
