package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/katsuokaisao/gin-play/domain"
)

func JWTMiddleware(parser domain.JWTParser, requiredScope domain.Scope) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
			c.Abort()
			return
		}

		bearer := "Bearer "
		if len(tokenString) <= len(bearer) || tokenString[:len(bearer)] != bearer {
			c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
			c.Abort()
			return
		}

		tokenString = tokenString[len(bearer):]
		claims, err := parser.Parse(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
			c.Abort()
			return
		}

		if !parser.HasScope(claims, requiredScope) {
			c.JSON(http.StatusForbidden, gin.H{"message": http.StatusText(http.StatusForbidden)})
			c.Abort()
			return
		}

		c.Set("claims", claims)
	}
}
