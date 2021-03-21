package controller

import (
	"awesomeProject2/pkg/internal"
	"awesomeProject2/pkg/internal/model"
	"awesomeProject2/pkg/internal/store"
	"awesomeProject2/pkg/internal/utils"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
	"strconv"
)

type UserController struct {
	UserStore 		store.IUserRepository
	SessionStore	sessions.Store
}

func NewUserController(s *internal.Server) *UserController {
	return &UserController{
		UserStore: s.Store.User(),
		SessionStore: s.SessionStore,
	}
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request)  {
	type request struct {
		User struct {
			Login    string `json:"login"`
			Phone    int64  `json:"phone"`
			Name     string `json:"name"`
			Password string `json:"password"`
		} `json:"user"`
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err)
		return
	}

	user := &model.User{
		Login:    req.User.Login,
		Phone:    req.User.Phone,
		Name:     req.User.Name,
		Password: req.User.Password,
	}

	if err := u.UserStore.Create(user); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err)
		return
	}

	user.Sanitize()
	utils.Respond(w, r, http.StatusCreated, user)
}

func (u *UserController) ShowUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err)
		return
	}

	user, err := u.UserStore.Find(id)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err)
		return
	}

	user.Sanitize()
	utils.Respond(w, r, http.StatusOK, user)
}

func (u *UserController) Authorize(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, errors.New("Uncorrected request"))
		return
	}

	user, err := u.UserStore.FindByLogin(req.Login)
	if err != nil || !user.ComparePassword(req.Password) {
		utils.RespondError(w, http.StatusUnauthorized, errors.New("Access denied"))
		return
	}

	session, err := u.SessionStore.Get(r, "token")
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	session.Values["user_id"] = user.ID
	if err := u.SessionStore.Save(r, w, session); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	utils.Respond(w, r, http.StatusOK, nil)
}
