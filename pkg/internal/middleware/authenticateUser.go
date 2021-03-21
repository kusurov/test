package middleware

import (
	"awesomeProject2/pkg/internal"
	"awesomeProject2/pkg/internal/model"
	"awesomeProject2/pkg/internal/utils"
	"context"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

func AuthenticateUser(s *internal.Server) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := s.SessionStore.Get(r, "token")
			if err != nil {
				utils.RespondError(w, http.StatusUnauthorized, errors.New("Not authenticated"))
				return
			}

			id, ok := session.Values["user_id"]
			if !ok {
				utils.RespondError(w, http.StatusUnauthorized, errors.New("Not authenticated"))
				return
			}

			user, err := s.Store.User().Find(id.(int64))
			if err != nil {
				utils.RespondError(w, http.StatusUnauthorized, errors.New("Not authenticated"))
				return
			}

			user.Sanitize()
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), utils.ContextKeyUser, user)))
		})
	}
}

func GetAuthUser(r *http.Request) *model.User {
	return r.Context().Value(utils.ContextKeyUser).(*model.User)
}