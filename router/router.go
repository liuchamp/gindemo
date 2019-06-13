package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"p1/controller"
	_ "p1/docs"
	"p1/middleware/jwt"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/login", Login)

	apiVersionOne := router.Group("/api/v1/")
	apiVersionOne.Use(jwt.Jwt())
	apiVersionOne.GET("user", controller.Info)
	return router

}
