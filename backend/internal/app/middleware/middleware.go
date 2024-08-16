package middleware

import (
	"net/http"
	"geomap/internal/app/logger"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/matoous/go-nanoid/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Middleware struct {
	log *logger.Logger
}

func New(log *logger.Logger) *Middleware {
	l := &logger.Logger{Logger: log.Named("[MIDDLEWARE]")}
	l.Debug("initialized package")
	return &Middleware{log: l}
}

func (m Middleware) ResponseRequestLogger(c *gin.Context) {

	correlationID := uuid.Must(8)

	// c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
	// c.Writer.Header().Set("X-Correlation-ID", correlationID)
	c.Set("X-Correlation-ID", correlationID) //можно записывать в хэдер

	m.log.Info("[REQUEST]",
		zap.String("correlation-id", correlationID),
		zap.String("clientIP", c.ClientIP()),
		zap.String("method", c.Request.Method),
		zap.String("path", c.Request.RequestURI),
		zap.String("agent", c.Request.UserAgent()),
	)

	t := time.Now()

	c.Next()

	latency := time.Since(t)
	statusCode := c.Writer.Status()
	fields := []zapcore.Field{
		zap.String("correlation-id", correlationID),
		zap.Int("statusCode", statusCode),
		zap.String("status", http.StatusText(c.Writer.Status())),
	}
	switch {
	case statusCode == 404:
		m.log.Warn("[RESPONSE]", fields...)
	case statusCode >= 400 && statusCode <= 499 && statusCode != 404:
		fields = append(fields,
			zap.String("latency", latency.String()),
			zap.String("errors", c.Errors.String()),
		)
		m.log.Warn("[RESPONSE]", fields...)
	case statusCode >= 500:
		fields = append(fields,
			zap.String("latency", latency.String()),
			zap.String("errors", c.Errors.String()),
		)
		m.log.Error("[RESPONSE]", fields...)

	default:
		fields = append(fields, zap.String("latency", latency.String()))
		if c.Errors.String() != "" {
			fields = append(fields, zap.String("errors", c.Errors.String()))
		}
		m.log.Info("[RESPONSE]", fields...)
	}
}
