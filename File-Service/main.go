package main

import (
	"fmt"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/File"
	"github.com/hoangphuc28/CoursesOnline/File-Service/config"
	"github.com/hoangphuc28/CoursesOnline/File-Service/internal/delivery/rpc"
	"github.com/hoangphuc28/CoursesOnline/File-Service/pkg/upload"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	env := os.Getenv("ENV")
	fileName := "config/config-local.yml"
	if env == "app" {
		fileName = "config/config-app.yml"
	}

	cf, err := config.LoadConfig(fileName)
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	cf.AWS.APIKey = apiKey
	cf.AWS.SecretKey = secretKey
	fmt.Println(cf.AWS.APIKey)
	fmt.Println(cf.AWS.SecretKey)

	lis, err := net.Listen("tcp", ":"+cf.App.Port)
	fmt.Println("Auth Svc on", cf.App.Port)
	s3 := upload.NewS3Provider(cf)
	firebase := upload.NewFireBaseProvider(cf)
	uploadFileProvider := upload.NewUploadFileProvider(s3, firebase)
	hdl := rpc.NewUploadHandler(uploadFileProvider, cf)
	grpcServer := grpc.NewServer(grpc.MaxMsgSize(10485760),
		grpc.MaxRecvMsgSize(10485760),
		grpc.MaxSendMsgSize(10485760))
	File.RegisterFileServiceServer(grpcServer, hdl)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

}
