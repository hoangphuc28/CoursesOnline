package client

import (
	"fmt"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/User"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/config"
	"google.golang.org/grpc"
)

// import (
//
//	"fmt"
//	"github.com/Zhoangp/User-Service/config"
//	"github.com/Zhoangp/User-Service/pb"
//	"google.golang.org/grpc"
//
// )
func InitUserServiceClient(c *config.Config) (User.UserServiceClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.OtherServices.UserUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
		return nil, err
	}
	return User.NewUserServiceClient(cc), nil
}
