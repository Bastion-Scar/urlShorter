package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		after := time.Since(start)
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		requestURI := c.Request.RequestURI
		path := c.Request.URL.Path

		logger.Info("Request Info",
			zap.Int("status", status),
			zap.String("clientIP", clientIP),
			zap.String("path", path),
			zap.String("requestURI", requestURI),
			zap.Duration("duration", after))
		if logger == nil {
			panic("logger not set")
		}
	}
}
