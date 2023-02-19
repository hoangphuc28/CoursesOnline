package userhttp

import (
	"context"
	"fmt"
	"github.com/Zhoangp/User-Service/pb"
	"github.com/Zhoangp/User-Service/pkg/client"
)

func (hdl *userHandler) GetProfileInstructor(ctx context.Context, req *pb.GetUserInformationRequest) (*pb.GetProfileInstructorResponse, error) {
	instructor, err := hdl.UC.GetInstructor(req.UserId, req.UserId)
	if err != nil {
		return &pb.GetProfileInstructorResponse{
			Error: HandleError(err),
		}, nil
	}
	paypalService, err := client.InitPaymentClient(hdl.Cf)
	if err != nil {
		return &pb.GetProfileInstructorResponse{Error: HandleError(err)}, err
	}
	res, err := paypalService.GetPaypal(ctx, &pb.GetPayalRequest{UserId: req.UserId})
	if err != nil {
		fmt.Println(err)

		return &pb.GetProfileInstructorResponse{Error: HandleError(err)}, err
	}
	return &pb.GetProfileInstructorResponse{
		AccountPaypal: res.Email,
		Website:       instructor.Information.Website,
		Linkedin:      instructor.Information.Linkedin,
		Youtube:       instructor.Information.Youtube,
		Bio:           instructor.Information.Bio,
	}, nil
}
