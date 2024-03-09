package http

import (
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Error"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Mail"
	"github.com/hoangphuc28/CoursesOnline/Mail-Service/internal/model"
	"github.com/hoangphuc28/CoursesOnline/Mail-Service/pkg/common"
)

type MailUsecase interface {
	SendEmail(email *model.Email, content *model.SendTokenContent) error
}
type mailHandler struct {
	uc MailUsecase
	Mail.UnimplementedMailServiceServer
}

func NewMailHandler(uc MailUsecase) *mailHandler {
	return &mailHandler{uc: uc}
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
