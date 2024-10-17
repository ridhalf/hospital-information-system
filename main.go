package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"hospital-information-system/app"
	"hospital-information-system/auth"
	"hospital-information-system/controller"
	"hospital-information-system/repository"
	"hospital-information-system/service"
	"os"
)

func main() {
	app.Env()
	db := app.NewDB()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	authJwt := auth.NewJwt()
	//authMiddleware := middleware.AuthMiddleware(authJwt, userService)

	userController := controller.NewUserController(userService, authJwt)

	router := gin.Default()
	//blocked by cors policy
	router.Use(cors.Default())
	//blocked by cors policy
	api := router.Group("/api/v1")
	api.POST("/users", userController.Register)
	api.GET("/users/:id", userController.FindById)
	err := router.Run(os.Getenv("DOMAIN"))
	if err != nil {
		panic(err)
	}
}
