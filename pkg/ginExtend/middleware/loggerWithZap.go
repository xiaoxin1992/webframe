package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
	"webframe/pkg/logger"
)

func ZapWithConfig(logger *logger.LogZap) gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		path := context.Request.URL.Path
		query := context.Request.URL.RawQuery
		context.Next()
		end := time.Now()
		latency := end.Sub(start)
		if len(context.Errors) > 0 {
			for _, e := range context.Errors.Errors() {
				logger.S("gin").Desugar().Error(e)
			}
		} else {
			fields := []zapcore.Field{
				zap.Int("status", context.Writer.Status()),
				zap.String("method", context.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", context.ClientIP()),
				zap.String("user-agent", context.Request.UserAgent()),
				zap.Duration("time", latency),
			}
			logger.S("gin").Desugar().Info(path, fields...)
		}
	}
}
