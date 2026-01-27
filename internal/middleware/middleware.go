package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	pkg "github.com/lokicodess/CatalogX/pkg/config"
)

func LogRequest(app *pkg.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			method = c.Request.Method
			uri    = c.Request.RequestURI
			ip     = c.Request.RemoteAddr
		)

		// Process the request
		start := time.Now()
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Determine log level based on status code
		status := c.Writer.Status()
		switch {
		case status >= 500:
			// Server error - ERROR level
			app.Logger.Error("request completed",
				slog.Int("status", status),
				slog.String("method", method),
				slog.String("uri", uri),
				slog.String("ip", ip),
				slog.Duration("latency", latency),
			)
		case status >= 400:
			// Client error - WARN level
			app.Logger.Warn("request completed",
				slog.Int("status", status),
				slog.String("method", method),
				slog.String("uri", uri),
				slog.String("ip", ip),
				slog.Duration("latency", latency),
			)
		default:
			// Success - INFO level
			app.Logger.Info("request completed",
				slog.Int("status", status),
				slog.String("method", method),
				slog.String("uri", uri),
				slog.String("ip", ip),
				slog.Duration("latency", latency),
			)
		}

	}
}
