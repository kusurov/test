package router

import (
	"kusurovAPI/internal/controller"
	"kusurovAPI/internal/server"
	"kusurovAPI/internal/server/middleware"
)

func HandleRouter(s *server.Server) {
	user := controller.NewUserController(s)
	category := controller.NewCategoryController(s)
	product := controller.NewProductController(s)
	catalog := controller.NewCatalogController(s)

	s.Router.Use(middleware.WriterHeaders)
	s.Router.Use(middleware.LoggerRequests(s.Logger))
	s.Router.HandleFunc("/auth", user.Authorize).Methods("POST")

	api := s.Router.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthenticateUser(s))
	api.HandleFunc("/users", middleware.CheckAdminPermission(user.CreateUser)).Methods("POST")
	api.HandleFunc("/users/{id}", middleware.CheckAdminPermission(user.ShowUser)).Methods("GET")

	api.HandleFunc("/categories", middleware.CheckAdminPermission(category.CreateCategory)).Methods("POST")
	api.HandleFunc("/categories", category.ShowAllCategories).Methods("GET")
	api.HandleFunc("/categories/{id}", category.ShowCategory).Methods("GET")
	api.HandleFunc("/categories/{id}/addProduct",
		middleware.CheckAdminPermission(product.CreateProduct)).Methods("POST")
	api.HandleFunc("/categories/{id}/disable",
		middleware.CheckAdminPermission(category.DisableCategory)).Methods("PUT")
	api.HandleFunc("/categories/{id}/enable",
		middleware.CheckAdminPermission(category.EnableCategory)).Methods("PUT")

	api.HandleFunc("/products/{id}", product.ShowProduct).Methods("GET")
	api.HandleFunc("/products/{id}/disable",
		middleware.CheckAdminPermission(product.DisableProduct)).Methods("PUT")
	api.HandleFunc("/products/{id}/enable",
		middleware.CheckAdminPermission(product.EnableProduct)).Methods("PUT")

	api.HandleFunc("/catalog", catalog.ShowCatalog).Methods("GET")
	api.HandleFunc("/catalog/{category_id}", catalog.ShowCatalogFromCategory).Methods("GET")
}
