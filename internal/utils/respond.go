package utils

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RespondError(w http.ResponseWriter, r *http.Request, code int, err error) {
	Respond(w, r, code, map[string]string{"error": err.Error()})
}

func Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)

	if data != nil {
		logger := r.Context().Value(ContextKeyLogger).(*logrus.Logger)

		logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"type":        "respond",
			"code":        code,
			"json":        data,
		}).Infof("Ending request %s %s", r.Method, r.RequestURI)

		json.NewEncoder(w).Encode(data)
	}
}
