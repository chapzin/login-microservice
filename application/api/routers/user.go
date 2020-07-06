package routers

import (
	"github.com/chapzin/login-microservice/application/api/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine, router *gin.RouterGroup) {
	user := router.Group("/user")
	user.GET("/", controllers.TestUser)
}
