package middleware

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"kusurovAPI/internal/model"
	"kusurovAPI/internal/server"
	"kusurovAPI/internal/utils"
	"net/http"
)

func AuthenticateUser(s *server.Server) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := s.SessionStore.Get(r, "token")
			if err != nil {
				utils.RespondError(w, r, http.StatusUnauthorized, errors.New("Not authenticated"))
				return
			}

			id, ok := session.Values["user_id"]
			if !ok {
				utils.RespondError(w, r, http.StatusUnauthorized, errors.New("Not authenticated"))
				return
			}

			user, err := s.Store.User.Find(id.(int64))
			if err != nil {
				utils.RespondError(w, r, http.StatusUnauthorized, errors.New("Not authenticated"))
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
