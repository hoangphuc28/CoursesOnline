package userhttp

import (
	"context"
	User "github.com/hoangphuc28/CoursesOnline-ProtoFile/User"
)

func (hdl *userHandler) GetProfile(ctx context.Context, req *User.GetUserInformationRequest) (*User.GetProfileResponse, error) {
	res, err := hdl.UC.GetProfile(req.UserId)
	if err != nil {
		return &User.GetProfileResponse{Error: HandleError(err)}, nil
	}
	return res, nil
}
