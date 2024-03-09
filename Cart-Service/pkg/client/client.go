package client

import (
	"fmt"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
	"github.com/hoangphuc28/CoursesOnline/Cart-Service/config"
	"google.golang.org/grpc"
)

func InitCourseServiceClient(c *config.Config) (Course.CourseServiceClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.OtherServices.CourseUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
		return nil, err
	}
	return Course.NewCourseServiceClient(cc), nil
}
