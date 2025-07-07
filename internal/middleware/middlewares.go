package middleware

import (
	"context"
	"net/http"
	"time"

	logger "github.com/mcandemir/bilinkat/internal/logger"
	utils "github.com/mcandemir/bilinkat/internal/utils"
)

func Logger(logger *logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			duration := time.Since(start)

			logger.Info(
				r.Context(),
				"HTTP request completed",
				"method", r.Method,
				"path", r.URL.Path,
				"duration_ms", duration.Milliseconds(),
				"remote_addr", r.RemoteAddr,
				"user_agent", r.UserAgent(),
			)
		})
	}
}

func Recoverer(logger *logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.Error(
						r.Context(),
						"Panic",
						"err", err,
					)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := utils.GenerateRequestID()
		w.Header().Set("X-Request-ID", requestId)
		ctx := context.WithValue(r.Context(), "request_id", requestId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RealIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
