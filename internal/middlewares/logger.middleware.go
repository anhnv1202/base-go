package middlewares

import (
	"time"

	"github.com/anhnv1202/base-go/global"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// Logger returns a gin.HandlerFunc that logs HTTP requests using the global zap logger
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate request ID if not present
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Set("request_id", requestID)

		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Get status code
		status := c.Writer.Status()

		// Get client IP
		clientIP := c.ClientIP()

		// Build full path with query string
		if raw != "" {
			path = path + "?" + raw
		}

		// Log the request
		fields := []zap.Field{
			zap.String("request_id", requestID),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.Int("status", status),
			zap.Duration("latency", latency),
			zap.String("client_ip", clientIP),
			zap.String("user_agent", c.Request.UserAgent()),
		}
		// Add error if present
		if len(c.Errors) > 0 {
			fields = append(fields, zap.String("error", c.Errors.String()))
		}

		// Log based on status code
		switch {
		case status >= 500:
			global.Logger.Error("HTTP Request", fields...)
		case status >= 400:
			global.Logger.Warn("HTTP Request", fields...)
		default:
			global.Logger.Info("HTTP Request", fields...)
		}
	}
}
