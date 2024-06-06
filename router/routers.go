package router

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 * gin.RouterGroup := router.Group("/api/v1")
}
