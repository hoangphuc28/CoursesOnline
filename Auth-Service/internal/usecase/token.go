package usecase

import "github.com/hoangphuc28/CoursesOnline/Auth-Service/pkg/utils"

func (uc *userUseCase) GetNewToken(refreshToken string) (*utils.Token, error) {
	data, err := utils.ValidateJWT(refreshToken, uc.cf)
	if err != nil {
		return nil, err
	}
	user, err := uc.userRepo.FindDataWithCondition(map[string]any{"email": data.Email})
	if err != nil {
		return nil, err
	}
	token, err := utils.GenerateToken(utils.TokenPayload{Email: user.Email, Role: user.Role}, uc.cf.Service.AccessTokenExpiredIn, uc.cf.Service.Secret)
	if err != nil {
		return nil, err
	}
	return token, err
}
