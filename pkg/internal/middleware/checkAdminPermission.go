package middleware

import (
	"awesomeProject2/pkg/internal/utils"
	"errors"
	"net/http"
)

func CheckAdminPermission(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := GetAuthUser(r)

		if !auth.HasAdminPermission() {
			utils.RespondError(w, r, http.StatusUnauthorized, errors.New("dont have permissions"))
			return
		}

		next.ServeHTTP(w, r)
	}
}
