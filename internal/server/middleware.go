package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/google/uuid"
)

func structuredLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		start := time.Now()
		defer func() {
			if ww.Status() >= 500 {
				slog.Error("server error", "method", r.Method, "path", r.URL.Path, "status", ww.Status(), "remote", r.RemoteAddr, "user-agent", r.UserAgent(), "time", time.Since(start), "request-id", r.Header.Get("X-Request-Id"), "internal-id", r.Context().Value("internal-id"))
			} else if ww.Status() >= 400 {
				slog.Warn("client error", "method", r.Method, "path", r.URL.Path, "status", ww.Status(), "remote", r.RemoteAddr, "user-agent", r.UserAgent(), "time", time.Since(start), "request-id", r.Header.Get("X-Request-Id"), "internal-id", r.Context().Value("internal-id"))
			} else if ww.Status() >= 300 {
				slog.Info("redirect", "method", r.Method, "path", r.URL.Path, "status", ww.Status(), "remote", r.RemoteAddr, "user-agent", r.UserAgent(), "time", time.Since(start), "request-id", r.Header.Get("X-Request-Id"), "internal-id", r.Context().Value("internal-id"))
			} else {
				slog.Info("success", "method", r.Method, "path", r.URL.Path, "status", ww.Status(), "remote", r.RemoteAddr, "user-agent", r.UserAgent(), "time", time.Since(start), "request-id", r.Header.Get("X-Request-Id"), "internal-id", r.Context().Value("internal-id"))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func requestTracking(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-Id")
		if requestID == "" {
			requestID = uuid.New().String()
			r.Header.Set("X-Request-Id", requestID)
		}
		internalID := uuid.New().String()
		r = r.WithContext(context.WithValue(r.Context(), "internal-id", internalID))
		next.ServeHTTP(w, r)
	})
}
