package routers

import "github.com/gin-gonic/gin"

func Initialize(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	UserRouter(r, v1)
}
