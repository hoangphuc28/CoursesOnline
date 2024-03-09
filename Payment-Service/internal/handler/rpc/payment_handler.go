package rpc

import (
	"context"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Cart"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Error"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Payment"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/config"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/pkg/client"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/pkg/common"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/pkg/utils"
)

type PaymentUseCase interface {
	CheckOutWithPaypal(cart *Cart.Cart, userId string) (*Payment.CheckoutResponse, error)
	CaptureWithPaypal(token string, orderId string) (*utils.TokenPayload, error)
	ConnectPaypal(token string, userId string) (string, error)
	GetPayment(req *Payment.GetPaymentRequest) (*Payment.GetPaymentResponse, error)
	GetPaypal(req *Payment.GetPayalRequest) (*Payment.GetPayalResponse, error)
	Refund(rq *Payment.RefundRequest) error
}
type paymentHandler struct {
	uc PaymentUseCase
	Payment.UnimplementedPaymentServiceServer
	cf *config.Config
}

func NewPaymentHandler(uc PaymentUseCase, cf *config.Config) *paymentHandler {
	return &paymentHandler{uc: uc, cf: cf}
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
func (hdl paymentHandler) CaptureWithPaypal(ctx context.Context, request *Payment.CaptureRequest) (*Payment.CaptureResponse, error) {
	payload, err := hdl.uc.CaptureWithPaypal(request.Token, request.OrderId)
	if err != nil {
		return &Payment.CaptureResponse{
			Error: HandleError(err),
		}, nil
	}
	cartClient, err := client.InitCartServiceClient(hdl.cf)
	if err != nil {
		return &Payment.CaptureResponse{
			Error: HandleError(err),
		}, nil
	}

	_, err = cartClient.ResetCart(ctx, &Cart.ResetCartRequest{
		CartId: payload.Cart.Id,
	})

	if err != nil {
		return &Payment.CaptureResponse{
			Error: HandleError(err),
		}, nil
	}

	courseClient, err := client.InitCourseServiceClient(hdl.cf)
	if err != nil {
		return &Payment.CaptureResponse{
			Error: HandleError(err),
		}, nil
	}
	for _, i := range payload.Cart.Courses {
		resCourse, err := courseClient.Enrollment(ctx, &Course.EnrollmentRequest{
			UserId:   payload.UserId,
			CourseId: i.Id,
		})
		if err != nil {
			return &Payment.CaptureResponse{
				Error: HandleError(err),
			}, nil
		}
		if resCourse.Error != nil {
			return &Payment.CaptureResponse{
				Error: resCourse.Error,
			}, nil
		}
	}
	return &Payment.CaptureResponse{}, nil
}
func (hdl paymentHandler) ConnectPaypalAccount(ctx context.Context, req *Payment.ConnectPaypalRequest) (*Payment.ConnectPaypalResponse, error) {
	email, err := hdl.uc.ConnectPaypal(req.IdentifyToken.Token, req.UserId)
	if err != nil {
		return &Payment.ConnectPaypalResponse{
			Error: HandleError(err),
		}, nil
	}
	return &Payment.ConnectPaypalResponse{
		Email: email,
	}, nil
}
