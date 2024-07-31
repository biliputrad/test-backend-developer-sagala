package route

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"test-backend-developer-sagala/common/constants"
	"test-backend-developer-sagala/config/env"
)

func InitRouter(config env.Config) *gin.Engine {
	if config.GinMode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch},
		AllowHeaders:    []string{"Content-Type", "Authorization", "Content-Length"},
	}))

	// limit file upload
	router.MaxMultipartMemory = 2 << 20

	return router
}

func RunRoute(config env.Config, router *gin.Engine) {
	err := router.Run(fmt.Sprintf(":%s", config.GinPort))
	if err != nil {
		message := fmt.Sprintf("%s failed to start server", constants.Server)
		log.Fatal(message)
	}
}
