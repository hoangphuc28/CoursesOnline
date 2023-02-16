package http

import (
	"errors"
	"fmt"
	"github.com/hoangphuc28/CoursesOnline/API-Gateway/config"
	"github.com/hoangphuc28/CoursesOnline/API-Gateway/pkg/common"
	"github.com/hoangphuc28/CoursesOnline/API-Gateway/services/auth/internal/model"
	pb "github.com/hoangphuc28/CoursesOnline/Proto/Auth-Service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	cf     *config.Config
	client pb.AuthServiceClient
}

func NewAuthHandler(cf *config.Config, client pb.AuthServiceClient) *AuthHandler {
	return &AuthHandler{cf, client}
}
func (hdl AuthHandler) GetTokenVerifyAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var rq model.GetTokenVerifyAccountRequest
		if err := ctx.ShouldBind(&rq); err != nil {
			panic(err)
		}
		if rq.Email == "" {
			panic(common.NewCustomError(errors.New("email is required"), "email is required"))
		}
		_, err := hdl.client.GetTokenVerifyAccount(ctx, &pb.VerifyAccountRequest{
			Email: rq.Email,
		})

		if err != nil {
			fmt.Println(err)

			panic(err)
		}

		ctx.JSON(200, gin.H{"Message": "Token has been sent to your email!"})
	}

}
func (hdl AuthHandler) GetTokenResetPassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var rq model.GetTokenVerifyAccountRequest
		if err := ctx.ShouldBind(&rq); err != nil {
			panic(err)
		}
		_, err := hdl.client.GetTokenResetPassword(ctx, &pb.VerifyAccountRequest{
			Email: rq.Email,
		})
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		ctx.JSON(200, gin.H{"Message": "Token has been sent to your email!"})
	}

}
func (hdl AuthHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		b := model.LoginRequest{}

		if err := ctx.BindJSON(&b); err != nil {
			panic(err)
		}
		res, err := hdl.client.Login(ctx, &pb.LoginRequest{
			Email:    b.Email,
			Password: b.Password,
		})
		if err != nil {
			panic(err)
		}
		ctx.SetCookie("refresh_token", res.RefreshToken, 3600*720, "/", "localhost", false, true)
		ctx.JSON(200, &res)

	}
}
func (hdl AuthHandler) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := model.RegisterRequest{}
		if err := ctx.BindJSON(&body); err != nil {
			panic(err)
		}
		_, err := hdl.client.Register(ctx, &pb.RegisterRequest{
			FirstName:   body.FirstName,
			LastName:    body.LastName,
			PhoneNumber: body.PhoneNumber,
			Email:       body.Email,
			Password:    body.Password,
			Address:     body.Address,
			Role:        body.Role,
		})

		if err != nil {
			panic(err)
		}

		ctx.JSON(200, gin.H{"Message": "Create User-Service successfully"})
	}

}
func (hdl AuthHandler) NewToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		refreshToken, err := ctx.Cookie("refresh_token")
		if err != nil {
			panic(err)
		}
		res, err := hdl.client.NewToken(ctx, &pb.NewTokenRequest{RefreshToken: refreshToken})
		if err != nil {

			panic(err)
		}
		ctx.JSON(200, &res)
	}

}
func (hdl AuthHandler) VerifyAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.MustGet("emailUser").(string)

		_, err := hdl.client.VerifyAccount(ctx, &pb.VerifyAccountRequest{
			Email: email,
		})

		if err != nil {
			panic(err)
		}

		ctx.JSON(200, gin.H{"Message": "Account has been verified"})
	}

}
