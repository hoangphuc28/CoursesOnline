package rpc

import (
	"context"
	"fmt"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Mail"
	"github.com/hoangphuc28/CoursesOnline/Mail-Service/internal/model"
)

func (hdl mailHandler) SendTokenVerifyAccount(context context.Context, request *Mail.SendTokenVerifyAccountRequest) (*Mail.SendTokenVerifyAccountResponse, error) {
	err := hdl.uc.SendEmail(&model.Email{
		DestMail: request.Mail.DestMail,
		Subject:  request.Mail.Subject,
	}, &model.SendTokenContent{
		Name: request.Name,
		Url:  request.Url + request.Token,
	})
	if err != nil {
		fmt.Println(err)
		return &Mail.SendTokenVerifyAccountResponse{
			Error: HandleError(err),
		}, nil
	}
	return &Mail.SendTokenVerifyAccountResponse{}, nil

}
