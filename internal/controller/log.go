package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

const (
	defaultMessage string = "Request log"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		url := c.Request.URL.String()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		message := c.GetString("message")

		if message == "" {
			message = defaultMessage
		}

		log.Printf("message: [%s] url: [%s] method: [%s] status: [%d]", message, url, method, statusCode)
	}
}
