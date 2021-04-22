package route

import (
	"golang-transaction/controller"
	"golang-transaction/middleware"
	"golang-transaction/repository"
	"golang-transaction/service"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//SetupRoutes : all the routes are defined here
func SetupRoutes(db *gorm.DB) {
	httpRouter := gin.Default()

	userRepository := repository.NewUserRepository(db)

	if err := userRepository.Migrate(); err != nil {
		log.Fatal("User migrate err", err)
	}
	userService := service.NewUserService(userRepository)

	userController := controller.NewUserController(userService)

	users := httpRouter.Group("users")

	users.GET("/", userController.GetAllUser)
	users.POST("/", userController.AddUser)

	httpRouter.POST("/money-transfer", middleware.DBTransactionMiddleware(db), userController.TransferMoney)
	httpRouter.Run()

}
