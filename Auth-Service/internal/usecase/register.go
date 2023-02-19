package usecase

import (
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/internal/model"
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/pkg/common"
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/pkg/utils"
)

func (uc *userUseCase) Register(data *model.Users) (*model.Users, string, error) {
	if user, _ := uc.userRepo.FindDataWithCondition(map[string]any{"email": data.Email}); user != nil {
		return nil, "", model.ErrEmailExisted
	}
	if err := data.PrepareCreate(); err != nil {
		return nil, "", err
	}
	data.Avatar = &common.Image{
		Id:     1,
		Url:    "https://firebasestorage.googleapis.com/v0/b/course-d9557.appspot.com/o/user%2Fta%CC%89i%20xuo%CC%82%CC%81ng.png?alt=media&token=2d69e8e9-20fc-4a2e-a597-b0b6096c25ce",
		Width:  "250px",
		Height: "250px",
	}

	if err := uc.userRepo.NewUsers(data); err != nil {
		return nil, "", err
	}
	token, err := utils.GenerateToken(utils.TokenPayload{Email: data.Email, Role: data.Role, Password: data.Password, Verified: false}, uc.cf.Service.ActiveTokenExpired, uc.cf.Service.Secret)
	if err != nil {
		return nil, "", err
	}
	data.FakeId = uc.h.Encode(data.Id)

	return data, token.AccessToken, nil
}
func (uc *userUseCase) GetTokenVerify(email string, key string) (*model.Users, string, error) {
	user, err := uc.userRepo.FindDataWithCondition(map[string]any{"email": email})
	if err != nil {
		return nil, "", model.ErrEmailOrPasswordInvalid
	}
	payload := utils.TokenPayload{Email: user.Email, Role: user.Role, Password: user.Password, Verified: user.Verified, Key: key}
	token, err := utils.GenerateToken(payload, uc.cf.Service.ActiveTokenExpired, uc.cf.Service.Secret)
	if err != nil {
		return nil, "", err
	}

	return user, token.AccessToken, nil
}
