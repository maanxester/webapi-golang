package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/maanxester/webapi-golang/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		users := main.Group("users")
		{
			users.GET("/:id", controllers.ShowUsers)
		}
	}
	return router
}


