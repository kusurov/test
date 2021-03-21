package utils

import (
	"encoding/json"
	"net/http"
)

func RespondError(w http.ResponseWriter, code int, err error)  {
	Respond(w, nil, code, map[string]string{"error": err.Error()})
}

func Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
