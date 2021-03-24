package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"kusurovAPI/internal/middleware"
	"kusurovAPI/internal/model"
	"kusurovAPI/internal/server"
	"kusurovAPI/internal/store"
	"kusurovAPI/internal/utils"
	"net/http"
	"strconv"
)

type ProductController struct {
	ProductStore  store.IProductRepository
	CategoryStore store.ICategoryRepository
}

func NewProductController(s *server.Server) *ProductController {
	return &ProductController{
		ProductStore:  s.Store.Product(),
		CategoryStore: s.Store.Category(),
	}
}

func (p *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	authUser := middleware.GetAuthUser(r)
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils.RespondError(w, r, http.StatusBadRequest, err)
		return
	}

	category, err := p.CategoryStore.Find(id, authUser)
	if err != nil {
		utils.RespondError(w, r, http.StatusBadRequest, err)
		return
	}

	type request struct {
		ID          int64   `json:"id"`
		Title       string  `json:"title"`
		Weight      float32 `json:"weight"`
		Size        float32 `json:"size"`
		Description string  `json:"description"`
		PhotoLink   string  `json:"photo_link"`
		Price       float32 `json:"price"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		utils.RespondError(w, r, http.StatusBadRequest, err)

		return
	}

	product := &model.Product{
		Title:       req.Title,
		Weight:      req.Weight,
		Size:        req.Size,
		Description: req.Description,
		PhotoLink:   req.PhotoLink,
		Price:       req.Price,
		Category:    *category,
	}

	if err := p.ProductStore.Create(product); err != nil {
		utils.RespondError(w, r, http.StatusBadRequest, err)
		return
	}

	utils.Respond(w, r, http.StatusCreated, product)
}

func (p *ProductController) ShowProduct(w http.ResponseWriter, r *http.Request) {
	authUser := middleware.GetAuthUser(r)
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils.RespondError(w, r, http.StatusBadRequest, err)
		return
	}

	product, err := p.ProductStore.Find(id, authUser)
	if err != nil {
		utils.RespondError(w, r, http.StatusBadRequest, err)
		return
	}

	utils.Respond(w, r, http.StatusOK, product)
}

func (p *ProductController) DisableProduct(w http.ResponseWriter, r *http.Request) {
	authUser := middleware.GetAuthUser(r)
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils.RespondError(w, r, http.StatusBadRequest, err)
		return
	}

	product, err := p.ProductStore.Find(id, authUser)
	if err != nil {
		utils.RespondError(w, r, http.StatusBadRequest, err)
		return
	}

	if product.Status == 0 {
		utils.Respond(w, r, http.StatusOK, product)
		return
	}

	updatedProduct, err := p.ProductStore.DisableProduct(product)
	if err != nil {
		utils.RespondError(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.Respond(w, r, http.StatusOK, updatedProduct)
}

func (p *ProductController) EnableProduct(w http.ResponseWriter, r *http.Request) {
	authUser := middleware.GetAuthUser(r)
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils.RespondError(w, r, http.StatusBadRequest, err)
		return
	}

	product, err := p.ProductStore.Find(id, authUser)
	if err != nil {
		utils.RespondError(w, r, http.StatusBadRequest, err)
		return
	}

	if product.Status == 1 {
		utils.Respond(w, r, http.StatusOK, product)
		return
	}

	updatedProduct, err := p.ProductStore.EnableProduct(product)
	if err != nil {
		utils.RespondError(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.Respond(w, r, http.StatusOK, updatedProduct)
}
