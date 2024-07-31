package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	controller "test-backend-developer-sagala/api/controller/user"
	service "test-backend-developer-sagala/application/user"
	repository "test-backend-developer-sagala/infrastructure/repository/user"
)

func UserRoute(db *gorm.DB, routerGroup *gin.RouterGroup) {
	// Repositories
	userRepo := repository.NewUserRepository(db)

	// Services
	userService := service.NewUserService(userRepo)

	// Controllers
	userController := controller.NewUserController(userService)

	// Endpoints
	routerGroup.POST("/user-service/", userController.Register)
	routerGroup.POST("/user-service/login", userController.Login)
}
