package router

import (
	"my-go-server/clinics"
	_ "my-go-server/docs"
	"my-go-server/users"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	users.RegisterUserRoutes(router)
	clinics.RegisterClinicRoutes(router)
	return router
}
