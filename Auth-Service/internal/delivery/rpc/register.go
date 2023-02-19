package userhttp

import (
	"context"
	"fmt"
	pb "github.com/hoangphuc28/CoursesOnline-ProtoFile/Auth"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Cart"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Mail"
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/internal/model"
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/pkg/client"
)

func (userHandler *UserHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	user := model.Users{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.PhoneNumber,
		Role:      req.Role,
		Address:   req.Address,
	}
	data, token, err := userHandler.UC.Register(&user)
	if err != nil {
		return &pb.RegisterResponse{
			Error: HandleError(err),
		}, nil
	}
	cartService, err := client.InitCartServiceClient(userHandler.cf)
	if err != nil {
		fmt.Println(err)
	}
	resCreateCart, err := cartService.CreateCart(ctx, &Cart.CreateCartRequest{
		UserId: data.FakeId,
	})
	fmt.Println(resCreateCart)
	if err != nil {
		fmt.Println(err)
	}
	mailService, err := client.InitServiceClient(userHandler.cf)
	if err != nil {
		fmt.Println(err)
		return &pb.RegisterResponse{
			Error: HandleError(err),
		}, nil
	}
	res, err := mailService.SendTokenVerifyAccount(ctx, &Mail.SendTokenVerifyAccountRequest{
		Mail: &Mail.Mail{
			DestMail: data.Email,
			Subject:  "Verify Account",
		},
		Token: token,
		Name:  data.LastName + " " + data.FirstName,
		Url:   "http://" + userHandler.cf.ClientSide.URL + "/register/successverify?token=",
	})

	if err != nil {
		fmt.Println(err)
		return &pb.RegisterResponse{
			Error: HandleError(err),
		}, nil
	}
	if res.Error != nil {

		return &pb.RegisterResponse{
			Error: res.Error,
		}, nil
	}

	return &pb.RegisterResponse{}, nil
}
