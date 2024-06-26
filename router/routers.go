package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffersonborges/gopportunities/docs"
	"github.com/jeffersonborges/gopportunities/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHanlder)
		v1.POST("/opening", handler.CreateOpeningHanlder)
		v1.DELETE("/opening", handler.DeleteOpeningHanlder)
		v1.PUT("/opening", handler.UpdateOpeningHanlder)
		v1.GET("/openings", handler.LitsOpeningsHanlder)
	}
	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
