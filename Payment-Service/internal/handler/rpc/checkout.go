package rpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Cart"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Payment"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/pkg/client"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/pkg/common"
	"strconv"
)

func (hdl paymentHandler) CheckOutWithPaypal(ctx context.Context, req *Payment.CheckoutRequest) (*Payment.CheckoutResponse, error) {
	cartClient, err := client.InitCartServiceClient(hdl.cf)
	if err != nil {
		return nil, err
	}
	res, err := cartClient.GetCart(ctx, &Cart.GetCartRequest{
		Id: req.UserId,
	})
	fmt.Println(res)
	total, err := strconv.Atoi(res.TotalCourse)
	if err != nil {
		return &Payment.CheckoutResponse{
			Error: HandleError(err),
		}, nil
	}
	if total == 0 {
		return &Payment.CheckoutResponse{
			Error: HandleError(common.NewCustomError(errors.New("\"Cart is empty\""), "Cart is empty")),
		}, nil
	}
	if err != nil {
		fmt.Println(err)
		return &Payment.CheckoutResponse{
			Error: HandleError(err),
		}, nil
	}
	response, err := hdl.uc.CheckOutWithPaypal(res.Cart, req.UserId)
	if err != nil {
		return nil, err
	}
	return response, nil
}
