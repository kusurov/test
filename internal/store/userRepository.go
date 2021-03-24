package store

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"kusurovAPI/internal/model"
)

type userRepository struct {
	db     *sql.DB
	logger *logrus.Logger
}

func (r *userRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	password, err := u.EncryptPassword()
	if err != nil {
		r.logger.Warn(err)

		return err
	}

	res, err := r.db.Exec(
		"INSERT INTO users (login, password, phone, name) VALUES (?, ?, ?, ?)",
		u.Login,
		password,
		u.Phone,
		u.Name,
	)
	if err != nil {
		r.logger.Warn(err)

		return err
	}

	// Получаю айдишник записи из бд
	id, err := res.LastInsertId()
	if err != nil {
		r.logger.Warn(err)

		return err
	}

	u.ID = id
	u.Access = 0

	return nil
}

func (r *userRepository) Find(id int64) (*model.User, error) {
	u := &model.User{}

	err := r.db.QueryRow(
		"SELECT user_id, login, phone, name, password, access FROM users WHERE `user_id` = ?", id,
	).Scan(
		&u.ID,
		&u.Login,
		&u.Phone,
		&u.Name,
		&u.Password,
		&u.Access,
	)

	if err != nil {
		r.logger.Warn(err)

		return nil, err
	}

	return u, nil
}

func (r *userRepository) FindByLogin(login string) (*model.User, error) {
	u := &model.User{}

	err := r.db.QueryRow(`
		SELECT user_id, login, phone, name, password, access 
			FROM users WHERE login = ?`,
		login,
	).Scan(
		&u.ID,
		&u.Login,
		&u.Phone,
		&u.Name,
		&u.Password,
		&u.Access,
	)

	if err != nil {
		r.logger.Warn(err)

		return nil, err
	}

	return u, nil
}
