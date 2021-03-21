package store

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

type Store struct {
	db 		*sql.DB
	logger	*logrus.Logger

	userRepository	*userRepository
	categoryRepository *categoryRepository
	productRepository *productRepository
}

func New(db *sql.DB, logger *logrus.Logger) *Store {
	return &Store{
		db: db,
		logger: logger,
	}
}

func (s *Store) User() IUserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &userRepository{
		db: s.db,
		logger: s.logger,
	}

	return s.userRepository
}

func (s *Store) Category() ICategoryRepository {
	if s.categoryRepository != nil {
		return s.categoryRepository
	}

	s.categoryRepository = &categoryRepository{
		db: s.db,
		logger: s.logger,
	}

	return s.categoryRepository
}

func (s *Store) Product() IProductRepository {
	if s.productRepository != nil {
		return s.productRepository
	}

	s.productRepository = &productRepository{
		db: s.db,
		logger: s.logger,
	}

	return s.productRepository
}

