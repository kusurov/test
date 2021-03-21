package middleware

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

func LoggerRequests(l *logrus.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger := l.WithFields(logrus.Fields{
				"remote_addr": r.RemoteAddr,
			})

			logger.Infof("Started %s %s", r.Method, r.RequestURI)

			next.ServeHTTP(w, r)
		})
	}
}
