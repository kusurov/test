package store

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"kusurovAPI/internal/model"
)

type productRepository struct {
	db     *sql.DB
	logger *logrus.Logger
}

func (r *productRepository) Create(product *model.Product) error {
	if err := product.Validate(); err != nil {
		return err
	}

	res, err := r.db.Exec(`INSERT INTO products (
					  title,
					  weight,
					  size,
					  description,
					  photo_link,
					  price,      
                      category_id
                      ) VALUES (?,?,?,?,?,?,?)`,
		product.Title,
		product.Weight,
		product.Size,
		product.Description,
		product.PhotoLink,
		product.Price,
		product.Category.ID,
	)

	if err != nil {
		r.logger.Warn(err)

		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		r.logger.Warn(err)

		return err
	}

	product.ID = id

	return nil
}

func (r *productRepository) Find(id int64, requester *model.User) (*model.Product, error) {
	p := &model.Product{}

	err := r.db.QueryRow(`SELECT product_id, title, weight, size, description, photo_link, price, status 
			FROM products WHERE status >= ? AND product_id = ?`,
		requesterStatusToView(requester),
		id,
	).Scan(
		&p.ID,
		&p.Title,
		&p.Weight,
		&p.Size,
		&p.Description,
		&p.PhotoLink,
		&p.Price,
		&p.Status,
	)

	if err != nil {
		r.logger.Warn(err)

		return nil, err
	}

	return p, nil
}

func (r *productRepository) DisableProduct(product *model.Product) (*model.Product, error) {
	if product.Status == 0 {
		return product, nil
	}

	_, err := r.db.Exec(`UPDATE products SET status=0 WHERE product_id=?`, product.ID)
	if err != nil {
		return nil, err
	}

	product.Status = 0

	return product, nil
}

func (r *productRepository) EnableProduct(product *model.Product) (*model.Product, error) {
	if product.Status == 1 {
		return product, nil
	}

	_, err := r.db.Exec(`UPDATE products SET status=1 WHERE product_id=?`, product.ID)
	if err != nil {
		return nil, err
	}

	product.Status = 1

	return product, nil
}

func (r *productRepository) GetAllProducts(requester *model.User, s *model.ProductSearchCriteria) ([]*model.Product, error) {
	var query = `SELECT
		products.product_id,
			products.title,
			products.weight,
			products.size,
			products.description,
			products.photo_link,
			products.price,
			products.status,
			categories.category_id,
			categories.title
		FROM products
		LEFT JOIN categories
		ON products.category_id = categories.category_id
		WHERE products.status >= ?
			AND categories.status >= ?`

	var rows *sql.Rows
	var err error

	if s.SearchCriteria.Title == "" {
		rows, err = r.db.Query(query,
			requesterStatusToView(requester),
			requesterStatusToView(requester),
		)
	} else {
		rows, err = r.db.Query(query+" AND products.title LIKE ?",
			requesterStatusToView(requester),
			requesterStatusToView(requester),
			"%"+s.SearchCriteria.Title+"%",
		)
	}

	if err != nil {
		r.logger.Warn(err)

		return nil, err
	}
	defer rows.Close()

	categories, err := parseRowsToModelProduct(rows)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *productRepository) GetAllProductsFromCategory(requester *model.User, categoryID int64, s *model.ProductSearchCriteria) ([]*model.Product, error) {
	var query = `SELECT 
			products.product_id, 
			products.title, 
			products.weight, 
			products.size, 
			products.description,
			products.photo_link, 
			products.price, 
			products.status,
			categories.category_id,
			categories.title
		FROM products
		LEFT JOIN categories 
		    ON products.category_id = categories.category_id
		WHERE products.status >= ? 
		  AND categories.status >= ? 
		  AND products.category_id = ?`

	var rows *sql.Rows
	var err error

	if s.SearchCriteria.Title == "" {
		rows, err = r.db.Query(query,
			requesterStatusToView(requester),
			requesterStatusToView(requester),
			categoryID,
		)
	} else {
		rows, err = r.db.Query(query+" AND products.title LIKE ?",
			requesterStatusToView(requester),
			requesterStatusToView(requester),
			categoryID,
			"%"+s.SearchCriteria.Title+"%",
		)
	}

	if err != nil {
		r.logger.Warn(err)

		return nil, err
	}
	defer rows.Close()

	categories, err := parseRowsToModelProduct(rows)
	if err != nil {
		r.logger.Warn(err)

		return nil, err
	}

	return categories, nil
}

func parseRowsToModelProduct(rows *sql.Rows) ([]*model.Product, error) {
	products := make([]*model.Product, 0)

	for rows.Next() {
		p := &model.Product{}

		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Weight,
			&p.Size,
			&p.Description,
			&p.PhotoLink,
			&p.Price,
			&p.Status,
			&p.Category.ID,
			&p.Category.Title,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}
