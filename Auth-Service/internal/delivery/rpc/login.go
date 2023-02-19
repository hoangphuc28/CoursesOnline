package userhttp

import (
	"context"
	pb "github.com/hoangphuc28/CoursesOnline-ProtoFile/Auth"
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/internal/model"
)

func (userHandler *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	var data model.UserLogin
	data.Email = req.Email
	data.Password = req.Password
	token, refreshToken, user, err := userHandler.UC.Login(&data)
	if err != nil {
		return &pb.LoginResponse{
			Error: HandleError(err),
		}, nil
	}

	return &pb.LoginResponse{
		Information: &pb.User{
			Id:          user.FakeId,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Email:       user.Email,
			PhoneNumber: user.Phone,
			Address:     user.Address,
			Avatar: &pb.Picture{
				Url:    user.Avatar.Url,
				Width:  user.Avatar.Width,
				Height: user.Avatar.Height,
			},
		},
		AccessToken:  token.AccessToken,
		RefreshToken: refreshToken.AccessToken,
		ExpiresAt:    uint32(token.ExpiresAt),
		TokenType:    "Bearer",
	}, nil
}
