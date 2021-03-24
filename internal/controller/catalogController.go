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

type CatalogController struct {
	ProductStore repository.IProductRepository
}

func NewCatalogController(s *server.Server) *CatalogController {
	return &CatalogController{
		ProductStore: s.Store.Product,
	}
}

func (c *CatalogController) ShowCatalog(w http.ResponseWriter, r *http.Request) {
	authUser := middleware.GetAuthUser(r)

	searchCriteria := &model.ProductSearchCriteria{}
	_ = json.NewDecoder(r.Body).Decode(searchCriteria)

	products, err := c.ProductStore.GetAllProducts(authUser, searchCriteria)
	if err != nil {
		utils.RespondError(w, r, http.StatusInternalServerError, errors.New("internal server error"))
		return
	}

	if len(products) == 0 {
		utils.RespondError(w, r, http.StatusNotFound, errors.New("not found products"))
		return
	}

	utils.Respond(w, r, http.StatusOK, products)
}

func (c *CatalogController) ShowCatalogFromCategory(w http.ResponseWriter, r *http.Request) {
	authUser := middleware.GetAuthUser(r)
	vars := mux.Vars(r)

	categoryId, err := strconv.ParseInt(vars["category_id"], 10, 64)
	if err != nil {
		utils.RespondError(w, r, http.StatusBadRequest, err)
		return
	}

	searchCriteria := &model.ProductSearchCriteria{}
	_ = json.NewDecoder(r.Body).Decode(searchCriteria)

	products, err := c.ProductStore.GetAllProductsFromCategory(authUser, categoryId, searchCriteria)

	if len(products) == 0 {
		utils.RespondError(w, r, http.StatusNotFound, errors.New("not found products"))
		return
	}

	utils.Respond(w, r, http.StatusOK, products)
}
