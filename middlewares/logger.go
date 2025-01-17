package middlewares

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// logs incoming HTTP requests and responses
func LogrusMiddleware(next http.Handler) http.Handler {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{}) // JSON logging for better structure

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lrw := &logResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(lrw, r) // Call the next handler

		duration := time.Since(start)

		logger.WithFields(logrus.Fields{
			"method":      r.Method,
			"url":         r.URL.Path,
			"status_code": lrw.statusCode,
			"duration":    duration.Seconds(),
			"user_agent":  r.UserAgent(),
			"remote_addr": r.RemoteAddr,
		}).Info("Handled request")
	})
}

// wraps http.ResponseWriter to capture the status code
type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *logResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
