package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(jwt.MapClaims)
		userRole := user["role"].(string)
		if userRole != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "permission to access endpoint denied",
				"error":   "NOT_AUTHORIZED",
			})
			return
		}
		c.Next()
	}
}
