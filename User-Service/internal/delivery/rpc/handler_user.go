package userhttp

import (
	"context"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Error"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/File"
	pb "github.com/hoangphuc28/CoursesOnline-ProtoFile/User"
	"github.com/hoangphuc28/CoursesOnline/User-Service/config"
	"github.com/hoangphuc28/CoursesOnline/User-Service/internal/model"
	"github.com/hoangphuc28/CoursesOnline/User-Service/pkg/client"
	"github.com/hoangphuc28/CoursesOnline/User-Service/pkg/common"
)

type userHandler struct {
	UC UserUseCase
	pb.UnimplementedUserServiceServer
	Cf *config.Config
}

type UserUseCase interface {
	ChangeUser(data *model.Users) error
	ChangePassword(data *model.UserChangePassword) error
	SendToken(email string) error
	NewInstructor(data *model.Instructor, email string) error
	ChangeAvatar(data *model.Users) error
	GetInstructor(id, key string) (*pb.GetInstructorInformationResponse, error)
	GetProfile(userId string) (*pb.GetProfileResponse, error)
}

func NewUserHandler(userUC UserUseCase, cf *config.Config) *userHandler {
	return &userHandler{UC: userUC, Cf: cf}
}
func HandleError(err error) *Error.ErrorResponse {
	if errors, ok := err.(*common.AppError); ok {
		return &Error.ErrorResponse{
			Code:    int64(errors.StatusCode),
			Message: errors.Message,
		}
	}
	appErr := common.ErrInternal(err.(error))
	return &Error.ErrorResponse{
		Code:    int64(appErr.StatusCode),
		Message: appErr.Message,
	}
}

func (userHandler *userHandler) ChangePassword(ctx context.Context, request *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	data := model.UserChangePassword{
		Email:       request.Email,
		OldPassword: request.Password,
		NewPass:     request.NewPassword,
	}
	if err := userHandler.UC.ChangePassword(&data); err != nil {
		return &pb.ChangePasswordResponse{
			Error: HandleError(err),
		}, nil
	}
	return &pb.ChangePasswordResponse{}, nil
}

func (userHandler *userHandler) UpdateAvatar(ctx context.Context, req *pb.UpdateAvatarRequest) (*pb.UpdateAvatarResponse, error) {
	cli, err := client.InitServiceClient(userHandler.Cf)
	if err != nil {
		return &pb.UpdateAvatarResponse{
			Error: HandleError(err),
		}, nil
	}

	res, err := cli.UploadAvatar(ctx, &File.UploadAvatarRequest{
		File: &File.File{
			FileName: req.FileName,
			Size:     req.Size,
			Content:  req.Content,
			Folder:   req.Folder,
		},
		OldUrl: req.FileName,
	})
	if err != nil {
		return &pb.UpdateAvatarResponse{
			Error: HandleError(err),
		}, nil
	}
	if res.Error != nil {
		return &pb.UpdateAvatarResponse{
			Error: res.Error,
		}, nil
	}

	data := model.Users{
		Email: req.Email,
		Avatar: &common.Image{
			Id:     1,
			Url:    res.Url,
			Width:  req.Width,
			Height: req.Height,
		},
	}
	if err := userHandler.UC.ChangeAvatar(&data); err != nil {
		return &pb.UpdateAvatarResponse{
			Error: HandleError(err),
		}, nil
	}

	return &pb.UpdateAvatarResponse{
		Url: res.Url,
	}, nil

}

func (userHandler *userHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	data := model.Users{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.PhoneNumber,
		Address:   req.Address,
	}
	if err := userHandler.UC.ChangeUser(&data); err != nil {
		return &pb.UpdateUserResponse{
			Error: HandleError(err),
		}, nil
	}
	return &pb.UpdateUserResponse{}, nil

}
func (userHandler *userHandler) GetTokenResetPass(ctx context.Context, req *pb.GetTokenResetPassRequest) (*pb.GetTokenResetPassResponse, error) {

	err := userHandler.UC.SendToken(req.Email)
	if err != nil {
		return &pb.GetTokenResetPassResponse{
			Error: HandleError(err),
		}, nil
	}
	return &pb.GetTokenResetPassResponse{}, nil
}
func (userHandler *userHandler) NewInstructor(ctx context.Context, req *pb.NewInstructorRequest) (*pb.NewInstructorResponse, error) {
	if err := userHandler.UC.NewInstructor(&model.Instructor{
		Website:  req.Website,
		LinkedIn: req.Linkedin,
		Youtube:  req.Youtube,
		Bio:      req.Bio,
	}, req.Email); err != nil {
		return &pb.NewInstructorResponse{
			Error: HandleError(err),
		}, nil
	}
	return &pb.NewInstructorResponse{}, nil
}
func (userHandler *userHandler) GetInstructor(ctx context.Context, req *pb.GetInstructorInformationRequest) (*pb.GetInstructorInformationResponse, error) {
	res, err := userHandler.UC.GetInstructor(req.Id, req.Key)
	if err != nil {
		return &pb.GetInstructorInformationResponse{
			Error: HandleError(err),
		}, nil
	}
	if res.Error != nil {
		return &pb.GetInstructorInformationResponse{
			Error: res.Error,
		}, nil
	}
	return res, nil
}
