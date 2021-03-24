package repository

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"kusurovAPI/internal/model"
	"kusurovAPI/internal/repository/mysql"
)

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

type Repositories struct {
	User     IUserRepository
	Category ICategoryRepository
	Product  IProductRepository
}

func NewRepositories(db *sql.DB, logger *logrus.Logger) *Repositories {
	return &Repositories{
		User:     mysql.NewUserRepository(db, logger),
		Category: mysql.NewCategoryRepository(db, logger),
		Product:  mysql.NewProductRepository(db, logger),
	}
}
