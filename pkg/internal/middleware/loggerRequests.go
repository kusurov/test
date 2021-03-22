package middleware

import (
	"awesomeProject2/pkg/internal/utils"
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

func LoggerRequests(logger *logrus.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.WithFields(logrus.Fields{
				"remote_addr": r.RemoteAddr,
				"type":        "request",
			}).Infof("Started request %s %s", r.Method, r.RequestURI)

			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), utils.ContextKeyLogger, logger)))
		})
	}
}
