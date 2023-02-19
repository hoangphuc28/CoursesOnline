package usecase

import (
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/config"
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/internal/model"
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/pkg/common"
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/pkg/utils"
)

type UserRepository interface {
	NewUsers(data *model.Users) error
	FindDataWithCondition(conditions map[string]any) (*model.Users, error)
	UpdateUser(user *model.Users, newInformation map[string]any) error
}

type userUseCase struct {
	cf       *config.Config
	userRepo UserRepository
	h        *utils.Hasher
}

func NewUserUseCase(userRepo UserRepository, cf *config.Config, h *utils.Hasher) *userUseCase {
	return &userUseCase{cf, userRepo, h}
}

func (uc *userUseCase) GetUserNotVerified(email string) error {
	user, err := uc.userRepo.FindDataWithCondition(map[string]any{"email": email, "verified": 0})
	if err != nil {
		return common.ErrEntityNotFound("Email", err)
	}
	if err := uc.userRepo.UpdateUser(user, map[string]any{"verified": 1}); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
