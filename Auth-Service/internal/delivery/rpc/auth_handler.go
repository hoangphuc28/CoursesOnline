package userhttp

import (
	pb "github.com/hoangphuc28/CoursesOnline-ProtoFile/Auth"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Error"
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/config"
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/internal/model"
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/pkg/common"
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/pkg/utils"
)

type UserHandler struct {
	UC UserUseCase
	pb.UnimplementedAuthServiceServer
	cf *config.Config
}
type UserUseCase interface {
	Register(data *model.Users) (*model.Users, string, error)
	Login(data *model.UserLogin) (*utils.Token, *utils.Token, *model.Users, error)
	GetNewToken(refreshToken string) (*utils.Token, error)
	GetUserNotVerified(email string) error
	GetTokenVerify(email string, key string) (*model.Users, string, error)
}

func NewUserHandler(cf *config.Config, userUC UserUseCase) *UserHandler {
	return &UserHandler{cf: cf, UC: userUC}
}

func HandleError(e error) *Error.ErrorResponse {
	if errors, ok := e.(*common.AppError); ok {
		return &Error.ErrorResponse{
			Code:    int64(errors.StatusCode),
			Message: errors.Message,
		}
	}
	appErr := common.ErrInternal(e.(error))
	return &Error.ErrorResponse{
		Code:    int64(appErr.StatusCode),
		Message: appErr.Message,
	}
}
