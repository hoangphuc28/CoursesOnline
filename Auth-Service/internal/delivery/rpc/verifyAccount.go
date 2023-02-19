package userhttp

import (
	"context"
	pb "github.com/hoangphuc28/CoursesOnline-ProtoFile/Auth"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Mail"
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/pkg/client"
)

func (userHandler *UserHandler) VerifyAccount(ctx context.Context, req *pb.VerifyAccountRequest) (*pb.VerifyAccountResponse, error) {
	if err := userHandler.UC.GetUserNotVerified(req.Email); err != nil {
		return &pb.VerifyAccountResponse{
			Error: HandleError(err),
		}, nil
	}
	return &pb.VerifyAccountResponse{}, nil
}
func (userHandler *UserHandler) GetTokenVerifyAccount(ctx context.Context, req *pb.VerifyAccountRequest) (*pb.VerifyAccountResponse, error) {
	user, token, err := userHandler.UC.GetTokenVerify(req.Email, "verify")
	mailService, err := client.InitServiceClient(userHandler.cf)
	if err != nil {
		return &pb.VerifyAccountResponse{
			Error: HandleError(err),
		}, nil
	}
	res, err := mailService.SendTokenVerifyAccount(ctx, &Mail.SendTokenVerifyAccountRequest{
		Mail: &Mail.Mail{
			DestMail: user.Email,
			Subject:  "Verify Account",
		},
		Token: token,
		Name:  user.FirstName + " " + user.LastName,
		Url:   "http://" + userHandler.cf.ClientSide.URL + "/register/successverify?token=",
	})

	if err != nil {
		return &pb.VerifyAccountResponse{
			Error: HandleError(err),
		}, nil
	}
	if res.Error != nil {
		return &pb.VerifyAccountResponse{
			Error: res.Error,
		}, nil
	}
	return &pb.VerifyAccountResponse{}, nil
}
