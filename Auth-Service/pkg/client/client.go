package client

import (
	"fmt"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Cart"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Mail"
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/config"

	"google.golang.org/grpc"
)

func InitServiceClient(c *config.Config) (Mail.MailServiceClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.OtherServices.MailServiceURL, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
		return nil, err
	}
	return Mail.NewMailServiceClient(cc), nil
}
func InitCartServiceClient(c *config.Config) (Cart.CartServiceClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.OtherServices.CartServiceUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
		return nil, err
	}
	fmt.Println(cc)
	return Cart.NewCartServiceClient(cc), nil
}
