package middleware

import (
	"net/http"

	"cospend/constant"
	"cospend/models"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a sample middleware for authentication and authorization using JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader(constant.AUTHORIZATION)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, models.Response{
				Code:         http.StatusUnauthorized,
				ResponseCode: constant.FAILED_AUTHORIZED,
				ResponseDesc: http.StatusText(http.StatusUnauthorized),
			})
			c.Abort()
			return
		}

		decodes, err := JwtClaim(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.Response{
				Code:         http.StatusUnauthorized,
				ResponseCode: constant.FAILED_AUTHORIZED,
				ResponseDesc: http.StatusText(http.StatusUnauthorized),
			})
			c.Abort()
			return
		}

		c.Set(constant.GIN_KEY, decodes)

		c.Next()
	}
}

// GetUserClaims возвращает пользовательские данные из JWT, сохранённые в контексте
func GetUserClaims(c *gin.Context) *models.UserClaims {
	claims, exists := c.Get(constant.GIN_KEY)
	if !exists {
		return nil
	}
	userClaims, ok := claims.(*models.UserClaims)
	if !ok {
		return nil
	}
	return userClaims
}