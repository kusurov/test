package store

import "awesomeProject2/pkg/internal/model"

type IUserRepository interface {
	Create(*model.User) error

	Find(int64) (*model.User, error)

	FindByLogin(string) (*model.User, error)
}

type ICategoryRepository interface {
	Create(*model.Category) error

	Find(int64, *model.User) (*model.Category, error)

	DisableCategory(*model.Category) (*model.Category, error)

	EnableCategory(*model.Category) (*model.Category, error)

	GetAllCategories(*model.User) ([]*model.Category, error)
}

type IProductRepository interface {
	Create(*model.Product) error

	Find(int64, *model.User) (*model.Product, error)

	DisableProduct(*model.Product) (*model.Product, error)

	EnableProduct(*model.Product) (*model.Product, error)

	GetAllProducts(*model.User, *model.ProductSearchCriteria) ([]*model.Product, error)

	GetAllProductsFromCategory(*model.User, int64, *model.ProductSearchCriteria) ([]*model.Product, error)
}
