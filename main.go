package main

import (
	"myproperty-clean-architecture/config"
	"myproperty-clean-architecture/controller"
	"myproperty-clean-architecture/repository"
	"myproperty-clean-architecture/service"

	"github.com/labstack/echo/v4"
)


func main() {
	conn := config.InitDB()
	e := echo.New()
	config.InitMigrate()


	// Setup Repository
	propertiesRepository := repository.NewPropertiesRepository(conn)


	// Setup Service
	propertiesService := service.NewPropertiesService(&propertiesRepository)

	// Setup Controller
	propertiesController := controller.NewPropertiesController(&propertiesService)

	// // Setup Fiber
	// app := fiber.New(config.NewFiberConfig())
	// app.Use(recover.New())

	// Setup Routing
	propertiesController.Route(e)

	// Start App
	e.Start(":8080")

}