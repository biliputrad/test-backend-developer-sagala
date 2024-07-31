package article

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	controller "test-backend-developer-sagala/api/controller/article"
	service "test-backend-developer-sagala/application/article"
	"test-backend-developer-sagala/config/database/paginate"
	jwtToken "test-backend-developer-sagala/config/jwt-token"
	repository "test-backend-developer-sagala/infrastructure/repository/article"
)

func ArticleRoute(db *gorm.DB, routerGroup *gin.RouterGroup) {
	// Repositories
	articleRepo := repository.NewArticleRepository(db)

	// Services
	articleService := service.NewArticleService(articleRepo)

	//paginate
	pagination := paginate.NewPagination()

	// Controllers
	articleController := controller.NewArticleController(articleService, *pagination)

	// Endpoints
	routerGroup.POST("/article-service/", jwtToken.Middleware(), articleController.Create)
	routerGroup.GET("/article-service/:id", jwtToken.Middleware(), articleController.FindById)
	routerGroup.GET("/article-service/", jwtToken.Middleware(), articleController.FindAll)
	routerGroup.GET("/article-service/without-pagination", jwtToken.Middleware(), articleController.FindAllWithoutPagination)
	routerGroup.DELETE("/article-service/:id", jwtToken.Middleware(), articleController.Delete)
}
