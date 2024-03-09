package client

import (
	"fmt"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Cart"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/config"
	"google.golang.org/grpc"
)

func InitCartServiceClient(c *config.Config) (Cart.CartServiceClient, error) {
	// using WithInsecure() because no SSL running

	cc, err := grpc.Dial(c.OtherServices.CartServiceUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
		return nil, err
	}
	return Cart.NewCartServiceClient(cc), nil
}
func InitCourseServiceClient(c *config.Config) (Course.CourseServiceClient, error) {
	// using WithInsecure() because no SSL running

	cc, err := grpc.Dial(c.OtherServices.CourseServiceUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
		return nil, err
	}
	return Course.NewCourseServiceClient(cc), nil
}
