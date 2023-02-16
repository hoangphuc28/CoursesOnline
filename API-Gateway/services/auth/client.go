package auth

import (
	"fmt"
	"github.com/hoangphuc28/CoursesOnline/API-Gateway/config"
	pb "github.com/hoangphuc28/CoursesOnline/Proto/Auth-Service"
	"google.golang.org/grpc"
)

// InitServiceClient func used to define auth service client
func NewAuthServiceClient(c *config.Config) pb.AuthServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.Services.AuthUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
