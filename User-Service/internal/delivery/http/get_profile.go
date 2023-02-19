package userhttp

import (
	"context"
	User_Service "github.com/hoangphuc28/CoursesOnline/Proto/v1.1/User-Service"
)

func (hdl *userHandler) GetProfile(ctx context.Context, req User_Service.GetUserInformationRequest) (*pb.GetProfileResponse, error) {
	res, err := hdl.UC.GetProfile(req.UserId)
	if err != nil {
		return &pb.GetProfileResponse{Error: HandleError(err)}, nil
	}
	return res, nil
}
