package usecase

import (
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/internal/model"
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/pkg/common"
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/pkg/utils"
)

func (uc *userUseCase) Login(data *model.UserLogin) (*utils.Token, *utils.Token, *model.Users, error) {
	user, err := uc.userRepo.FindDataWithCondition(map[string]any{"email": data.Email})
	if err != nil {
		return nil, nil, nil, model.ErrEmailOrPasswordInvalid
	}
	if !user.Verified {
		return nil, nil, nil, model.ErrAccountNotVerified
	}

	if err := utils.ComparePassword(user.Password, data.Password); err != nil {
		return nil, nil, nil, model.ErrEmailOrPasswordInvalid
	}
	user.FakeId = uc.h.Encode(user.Id)
	token, err := utils.GenerateToken(utils.TokenPayload{Id: user.FakeId, Email: user.Email, Role: user.Role, Password: user.Password}, uc.cf.Service.AccessTokenExpiredIn, uc.cf.Service.Secret)
	if err != nil {
		return nil, nil, nil, common.ErrInternal(err)
	}
	refreshToken, err := utils.GenerateToken(utils.TokenPayload{Id: user.FakeId, Email: user.Email, Role: user.Role, Password: user.Password}, uc.cf.Service.RefreshTokenExpiredIn, uc.cf.Service.Secret)
	if err != nil {
		return nil, nil, nil, common.ErrInternal(err)
	}
	if err := uc.userRepo.UpdateUser(user, map[string]any{"refresh_token": refreshToken.AccessToken}); err != nil {
		return nil, nil, nil, common.ErrInternal(err)
	}
	return token, refreshToken, user, nil
}
