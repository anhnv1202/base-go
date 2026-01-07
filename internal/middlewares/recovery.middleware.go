package middlewares

import (
	"net/http"
	"runtime/debug"

	"github.com/anhnv1202/base-go/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Recovery returns a gin.HandlerFunc that recovers from panics and logs them using the global zap logger
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Get request ID if available
				requestID, _ := c.Get("request_id")
				requestIDStr := ""
				if id, ok := requestID.(string); ok {
					requestIDStr = id
				}

				// Get stack trace
				stack := string(debug.Stack())

				// Log the panic with stack trace
				global.Logger.Error("Panic recovered",
					zap.String("request_id", requestIDStr),
					zap.String("method", c.Request.Method),
					zap.String("path", c.Request.URL.Path),
					zap.String("client_ip", c.ClientIP()),
					zap.Any("error", err),
					zap.String("stack", stack),
				)

				// Return 500 error response
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": gin.H{
						"code":       "INTERNAL_SERVER_ERROR",
						"message":    "Internal server error",
						"request_id": requestIDStr,
					},
				})

				// Abort the request
				c.Abort()
			}
		}()

		c.Next()
	}
}
