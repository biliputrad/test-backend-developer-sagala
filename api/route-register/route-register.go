package route_register

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"test-backend-developer-sagala/api/route/article"
	"test-backend-developer-sagala/api/route/user"
)

func RouteRegister(db *gorm.DB, routerGroup *gin.RouterGroup) {
	article.ArticleRoute(db, routerGroup)
	user.UserRoute(db, routerGroup)
}
