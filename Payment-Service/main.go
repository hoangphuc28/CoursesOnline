package main

import (
	"fmt"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Payment"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/config"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/internal/handler/rpc"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/internal/payment_methos/paypal"
	repo2 "github.com/hoangphuc28/CoursesOnline/Payment-Service/internal/repo"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/internal/usecase"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/pkg/database/mysql"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/pkg/utils"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	env := os.Getenv("ENV")
	pathConfig := "config/config-local.yml"
	if env == "app" {
		pathConfig = "config/config-app.yml"
	}
	cf, err := config.LoadConfig(pathConfig)
	if err != nil {
		fmt.Println("Could not load config!")
		return
	}
	gormDb, err := mysql.NewMysql(cf)
	if err != nil {
		fmt.Println(err)
		return
	}

	lis, err := net.Listen("tcp", ":"+cf.Service.Port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Auth Svc on", cf.Service.Port)
	pp := paypal.NewPayPalHandler(cf)

	hasher := utils.NewHasher("courses", 3)

	repo := repo2.NewPaymentRepo(gormDb)
	uc := usecase.NewPaymentUseCase(repo, cf, pp, hasher)
	handler := rpc.NewPaymentHandler(uc, cf)
	grpcServer := grpc.NewServer()
	Payment.RegisterPaymentServiceServer(grpcServer, handler)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

}
