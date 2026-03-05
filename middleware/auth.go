package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidarHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		var clave = "sports-secure-2026"

		header := c.GetHeader("X-API-KEY")

		if header != clave {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
