package rpc

import (
	"context"
	"fmt"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Cart"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Error"
	"github.com/hoangphuc28/CoursesOnline/Cart-Service/config"
	"github.com/hoangphuc28/CoursesOnline/Cart-Service/pkg/client"
	"github.com/hoangphuc28/CoursesOnline/Cart-Service/pkg/common"
)

type CartUsecase interface {
	GetCart(fakeId string) (*Cart.GetCartResponse, error)
	AddToCart(cartId string, courseId string) error
	RemoveItem(cartId string, courseId string) error
	ResetCart(cartId string) error
	NewCart(userId string) error
}
type cartHandler struct {
	uc CartUsecase
	Cart.UnimplementedCartServiceServer
	cf *config.Config
}

func NewCartHandler(uc CartUsecase, cf *config.Config) *cartHandler {
	return &cartHandler{uc: uc, cf: cf}
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
func (hdl cartHandler) GetCart(ctx context.Context, request *Cart.GetCartRequest) (*Cart.GetCartResponse, error) {
	cart, err := hdl.uc.GetCart(request.Id)
	if err != nil {
		return &Cart.GetCartResponse{
			Error: HandleError(err),
		}, nil
	}
	return cart, nil
}
func (hdl cartHandler) AddItem(ctx context.Context, request *Cart.CartItemRequest) (*Cart.CartItemResponse, error) {
	courseService, err := client.InitCourseServiceClient(hdl.cf)
	if err != nil {
		fmt.Println(err)
		return &Cart.CartItemResponse{
			Error: HandleError(err),
		}, nil
	}

	course, err := courseService.GetCourse(ctx, &Course.GetCourseRequest{
		Id: request.CourseId,
	})

	if err != nil {
		fmt.Println(err)

		return &Cart.CartItemResponse{
			Error: HandleError(err),
		}, nil
	}
	if course.Error != nil {
		return &Cart.CartItemResponse{
			Error: HandleError(err),
		}, nil
	}

	if course.Course.Instructor.Id == request.CartId {
		fmt.Println(1)
		return &Cart.CartItemResponse{
			Error: HandleError(common.NewCustomError(err, "Your are an instructor of this course!")),
		}, nil
	}

	enrollments, err := courseService.GetEnrollments(ctx, &Course.GetEnrollmentsRequest{
		UserId: request.CartId,
	})

	for _, item := range enrollments.Enrollments {
		if request.CourseId == item.CourseId {
			return &Cart.CartItemResponse{
				Error: HandleError(common.NewCustomError(err, "You have been in this course!")),
			}, nil
		}
	}
	if err := hdl.uc.AddToCart(request.CartId, request.CourseId); err != nil {
		return &Cart.CartItemResponse{
			Error: HandleError(err),
		}, nil
	}
	return &Cart.CartItemResponse{}, nil
}
func (hdl cartHandler) RemoveItem(ctx context.Context, request *Cart.CartItemRequest) (*Cart.CartItemResponse, error) {
	if err := hdl.uc.RemoveItem(request.CartId, request.CourseId); err != nil {
		return &Cart.CartItemResponse{
			Error: HandleError(err),
		}, nil
	}
	return &Cart.CartItemResponse{}, nil
}
func (hdl cartHandler) ResetCart(ctx context.Context, request *Cart.ResetCartRequest) (*Cart.ResetCartResponse, error) {
	if err := hdl.uc.ResetCart(request.CartId); err != nil {
		return &Cart.ResetCartResponse{
			Error: HandleError(err),
		}, nil
	}
	return &Cart.ResetCartResponse{}, nil
}
func (hdl cartHandler) CreateCart(ctx context.Context, request *Cart.CreateCartRequest) (*Cart.CreateCartResponse, error) {

	if err := hdl.uc.NewCart(request.UserId); err != nil {
		return &Cart.CreateCartResponse{
			Error: HandleError(err),
		}, nil
	}
	return &Cart.CreateCartResponse{}, nil
}
