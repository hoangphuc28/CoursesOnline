package client

import (
	"fmt"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/File"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Payment"
	"github.com/hoangphuc28/CoursesOnline/User-Service/config"
	"google.golang.org/grpc"
)

func InitServiceClient(c *config.Config) (File.FileServiceClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.OtherServices.FileUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
		return nil, err
	}
	return File.NewFileServiceClient(cc), nil
}
func InitPaymentClient(c *config.Config) (Payment.PaymentServiceClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.OtherServices.PaymentUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
		return nil, err
	}
	return Payment.NewPaymentServiceClient(cc), nil
}
