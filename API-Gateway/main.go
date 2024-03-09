package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hoangphuc28/CoursesOnline/API-Gateway/config"
	"github.com/hoangphuc28/CoursesOnline/API-Gateway/middleware"
	"github.com/hoangphuc28/CoursesOnline/API-Gateway/pkg/utils"
	"github.com/hoangphuc28/CoursesOnline/API-Gateway/services/auth"
	"log"
	"os"
)

func main() {
	//Get environment variable
	env := os.Getenv("ENV")
	//Load config
	fileName := "config/config-local.yml"
	if env == "app" {
		fileName = "config/config-app.yml"
	}
	cf, err := config.LoadConfig(fileName)
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	if env == "app" {
		utils.RunDBMigration(cf)
	}
	//create middleware manager instance
	mdware := middleware.NewMiddlewareManager(cf)

	//config cors middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"authorization", "content-type"}
	config.AllowMethods = []string{"*"}
	config.AllowFiles = true

	//new instance of the Gin web framework
	r := gin.Default()

	//apply middleware for all requests
	r.Use(cors.New(config), mdware.Recover())

	//define routers for each service
	auth.RegisterAuthRoutes(r, cf, mdware)
	//user.RegisterUserRoutes(r, cf, mdware)
	//file_service.RegisterFileRoute(r, cf, mdware)
	//course.RegisterCourseService(r, cf, mdware)
	//cart.NewCartRoutes(r, cf, mdware)
	//payment.NewPaymentRoutes(r, cf, mdware)

	r.Run(cf.Services.Port)
}
