package routeres

import (
	_ "my-go-server/docs"
	"github.com/gin-gonic/gin"
	 "my-go-server/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(userService *services.UserService) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	RegisterUserRoutes(router, userService)
	RegisterClinicRoutes(router)
	return router
}
