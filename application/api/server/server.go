package server

import (
	middleware "github.com/chapzin/login-microservice/application/api/middlewares"
	"github.com/chapzin/login-microservice/application/api/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"
)

func Setup(db *gorm.DB, conn *amqp.Connection) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middleware.SetDBtoContext(db))
	r.Use(middleware.SetRabbitMQContext(conn))
	routers.Initialize(r)
	return r
}
