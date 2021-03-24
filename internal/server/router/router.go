package router

import (
	"kusurovAPI/internal/controller"
	"kusurovAPI/internal/server"
	mw "kusurovAPI/internal/server/middleware"
)

func HandleRouter(s *server.Server) {
	user := controller.NewUserController(s)
	category := controller.NewCategoryController(s)
	product := controller.NewProductController(s)
	catalog := controller.NewCatalogController(s)

	s.Router.Use(mw.WriterHeaders)
	s.Router.Use(mw.LoggerRequests(s.Logger))
	s.Router.HandleFunc("/auth", user.Authorize).Methods("POST")

	api := s.Router.PathPrefix("/api").Subrouter()
	api.Use(mw.AuthenticateUser(s))
	api.HandleFunc("/users", mw.CheckAdminPermission(user.CreateUser)).Methods("POST")
	api.HandleFunc("/users/{id}", mw.CheckAdminPermission(user.ShowUser)).Methods("GET")

	api.HandleFunc("/categories", mw.CheckAdminPermission(category.CreateCategory)).Methods("POST")
	api.HandleFunc("/categories", category.ShowAllCategories).Methods("GET")
	api.HandleFunc("/categories/{id}", category.ShowCategory).Methods("GET")
	api.HandleFunc("/categories/{id}/addProduct",
		mw.CheckAdminPermission(product.CreateProduct)).Methods("POST")
	api.HandleFunc("/categories/{id}/disable",
		mw.CheckAdminPermission(category.DisableCategory)).Methods("PUT")
	api.HandleFunc("/categories/{id}/enable",
		mw.CheckAdminPermission(category.EnableCategory)).Methods("PUT")

	api.HandleFunc("/products/{id}", product.ShowProduct).Methods("GET")
	api.HandleFunc("/products/{id}/disable",
		mw.CheckAdminPermission(product.DisableProduct)).Methods("PUT")
	api.HandleFunc("/products/{id}/enable",
		mw.CheckAdminPermission(product.EnableProduct)).Methods("PUT")

	api.HandleFunc("/catalog", catalog.ShowCatalog).Methods("GET")
	api.HandleFunc("/catalog/{category_id}", catalog.ShowCatalogFromCategory).Methods("GET")
}
