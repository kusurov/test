package controller

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"kusurovAPI/internal/model"
	"kusurovAPI/internal/repository"
	"kusurovAPI/internal/server"
	"kusurovAPI/internal/server/middleware"
	"kusurovAPI/internal/utils"
	"net/http"
	"strconv"
)

type CategoryController struct {
	CategoryStore repository.ICategoryRepository
}

func NewCategoryController(s *server.Server) *CategoryController {
	return &CategoryController{
		CategoryStore: s.Store.Category,
	}
}

func (c *CategoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Category struct {
			Title string `json:"title"`
		} `json:"category"`
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		utils.RespondError(w, r, http.StatusBadRequest, err)
		return
	}

	category := &model.Category{
		Title: req.Category.Title,
	}

	if err := c.CategoryStore.Create(category); err != nil {
		utils.RespondError(w, r, http.StatusBadRequest, err)
		return
	}

	utils.Respond(w, r, http.StatusCreated, category)
}

func (c *CategoryController) ShowCategory(w http.ResponseWriter, r *http.Request) {
	authUser := middleware.GetAuthUser(r)
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils.RespondError(w, r, http.StatusBadRequest, err)
		return
	}

	category, err := c.CategoryStore.Find(id, authUser)
	if err != nil {
		utils.RespondError(w, r, http.StatusBadRequest, err)
		return
	}

	utils.Respond(w, r, http.StatusOK, category)
}

func (c *CategoryController) ShowAllCategories(w http.ResponseWriter, r *http.Request) {
	authUser := middleware.GetAuthUser(r)

	categories, err := c.CategoryStore.GetAllCategories(authUser)
	if err != nil {
		utils.RespondError(w, r, http.StatusInternalServerError, errors.New("internal server error"))
		return
	}

	if len(categories) == 0 {
		utils.RespondError(w, r, http.StatusNotFound, errors.New("not found categories"))
		return
	}

	utils.Respond(w, r, http.StatusOK, categories)
}

func (c *CategoryController) EnableCategory(w http.ResponseWriter, r *http.Request) {
	authUser := middleware.GetAuthUser(r)
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils.RespondError(w, r, http.StatusBadRequest, err)
		return
	}

	category, err := c.CategoryStore.Find(id, authUser)
	if err != nil {
		utils.RespondError(w, r, http.StatusBadRequest, err)
		return
	}

	updatedCategory, err := c.CategoryStore.EnableCategory(category)
	if err != nil {
		utils.RespondError(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.Respond(w, r, http.StatusOK, updatedCategory)
}

func (c *CategoryController) DisableCategory(w http.ResponseWriter, r *http.Request) {
	authUser := middleware.GetAuthUser(r)
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils.RespondError(w, r, http.StatusBadRequest, err)
		return
	}

	category, err := c.CategoryStore.Find(id, authUser)
	if err != nil {
		utils.RespondError(w, r, http.StatusBadRequest, err)
		return
	}

	updatedCategory, err := c.CategoryStore.DisableCategory(category)
	if err != nil {
		utils.RespondError(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.Respond(w, r, http.StatusOK, updatedCategory)
}
