package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"
)

func SetDBtoContext(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}

func SetRabbitMQContext(conn *amqp.Connection) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("AMQP", conn)
		c.Next()
	}
}
