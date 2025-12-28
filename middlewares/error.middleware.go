package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			c.JSON(http.StatusInternalServerError, gin.H{
				"status":     false,
				"statusCode": 500,
				"message":    err.Error(),
			})
		}
	}
}
