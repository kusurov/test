package store

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"kusurovAPI/internal/model"
)

type categoryRepository struct {
	db     *sql.DB
	logger *logrus.Logger
}

func (r *categoryRepository) Create(c *model.Category) error {
	if err := c.Validate(); err != nil {
		return err
	}

	res, err := r.db.Exec("INSERT INTO categories (title) VALUES (?)", c.Title)
	if err != nil {
		r.logger.Warn(err)

		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		r.logger.Warn(err)

		return err
	}

	c.ID = id

	return nil
}

func (r *categoryRepository) Find(id int64, requester *model.User) (*model.Category, error) {
	c := &model.Category{}

	err := r.db.QueryRow(
		"SELECT category_id, title, status FROM categories WHERE `status` >= ? AND `category_id` = ?",
		requesterStatusToView(requester),
		id,
	).Scan(
		&c.ID,
		&c.Title,
		&c.Status,
	)

	if err != nil {
		r.logger.Warn(err)

		return nil, err
	}

	return c, nil
}

func (r *categoryRepository) DisableCategory(category *model.Category) (*model.Category, error) {
	if category.Status == 0 {
		return category, nil
	}

	_, err := r.db.Exec(`UPDATE categories SET status=0 WHERE category_id=?`, category.ID)
	if err != nil {
		return nil, err
	}

	category.Status = 0

	return category, nil
}

func (r *categoryRepository) EnableCategory(category *model.Category) (*model.Category, error) {
	if category.Status == 1 {
		return category, nil
	}

	_, err := r.db.Exec(`UPDATE categories SET status=1 WHERE category_id=?`, category.ID)
	if err != nil {
		return nil, err
	}

	category.Status = 1

	return category, nil
}

func (r *categoryRepository) GetAllCategories(requester *model.User) ([]*model.Category, error) {
	categories := make([]*model.Category, 0)

	rows, err := r.db.Query("SELECT category_id, title FROM categories WHERE `status` >= ?",
		requesterStatusToView(requester),
	)
	if err != nil {
		r.logger.Warn(err)

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		category := &model.Category{}

		err := rows.Scan(
			&category.ID,
			&category.Title,
		)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

//админ может просматривать заявки со статусом 0, т.е. заблокированные
func requesterStatusToView(requester *model.User) uint8 {
	var moreThanCategoryStatus uint8 = 1

	if requester.HasAdminPermission() {
		moreThanCategoryStatus = 0
	}

	return moreThanCategoryStatus
}
