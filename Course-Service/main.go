package main

import (
	"fmt"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/config"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/internal/delivery/https"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/internal/repo"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/internal/usecase"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/pkg/database/mysql"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/pkg/utils"
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
	hasher := utils.NewHasher("courses", 3)
	gormDb, err := mysql.NewMysql(cf)
	if err != nil {
		fmt.Println(err)
		return
	}
	courseRepo := repo.NewCoursesRepository(gormDb)
	courseUsecase := usecase.NewCoursesUseCase(courseRepo, hasher, cf)
	courseHandler := https.NewCoursesHandler(courseUsecase, cf)
	lis, err := net.Listen("tcp", ":"+cf.Service.Port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Course Svc on", cf.Service.Port)
	grpcServer := grpc.NewServer(grpc.MaxMsgSize(cf.Service.MaxSizeMess),
		grpc.MaxRecvMsgSize(cf.Service.MaxSizeMess),
		grpc.MaxSendMsgSize(cf.Service.MaxSizeMess))
	Course.RegisterCourseServiceServer(grpcServer, courseHandler)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
