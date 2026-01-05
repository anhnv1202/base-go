package middlewares

import (
	"github.com/anhnv1202/base-go/pkg/apperror"
	"github.com/anhnv1202/base-go/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return response.Wrap(func(c *gin.Context) (any, error) {
		token := c.GetHeader("Authorization")
		if token == "" {
			return nil, apperror.ErrUnauthorized("Authorization header is required")
		}
		c.Next()
		return nil, nil
	})
}
