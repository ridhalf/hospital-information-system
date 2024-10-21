package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"hospital-information-system/app"
	"hospital-information-system/auth"
	"hospital-information-system/controller"
	"hospital-information-system/middleware"
	"hospital-information-system/repository"
	"hospital-information-system/service"
	"os"
)

func main() {
	app.Env()
	db := app.NewDB()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	patientRepository := repository.NewPatientRepository(db)
	patientService := service.NewPatientService(patientRepository, userRepository)
	medicalRecord := repository.NewMedicalRecordRepository(db)
	medicalRecordService := service.NewMedicalRecordService(medicalRecord)

	authJwt := auth.NewJwt()
	authMiddleware := middleware.AuthMiddleware(authJwt, userService, patientService)

	userController := controller.NewUserController(userService, authJwt)
	patientController := controller.NewPatientController(patientService, authJwt)
	medicalRecordController := controller.NewMedicalRecordController(medicalRecordService, authJwt)

	router := gin.Default()
	//blocked by cors policy
	router.Use(cors.Default())
	//blocked by cors policy
	api := router.Group("/api/v1")
	api.POST("/users", userController.Register)
	api.POST("/users/login", userController.Login)
	api.GET("/users/:id", userController.FindById)

	api.POST("/patients/register", patientController.RegisterPatient)
	api.GET("/patients/:id", authMiddleware, patientController.FindById)

	api.GET("/medical_record/:patient_id", authMiddleware, medicalRecordController.FindByPatientID)
	err := router.Run(os.Getenv("DOMAIN"))
	if err != nil {
		panic(err)
	}
}
