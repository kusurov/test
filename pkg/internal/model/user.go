package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID 			int64 	`json:"id"`
	Login 		string 	`json:"login"`
	Phone	 	int64 	`json:"phone"`
	Name 		string 	`json:"name"`
	Password	string 	`json:"password,omitempty"`
	Access		int8	`json:"-"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Login, validation.Required),
		validation.Field(&u.Phone, validation.Required),
		validation.Field(&u.Name, validation.Required),
		validation.Field(&u.Password, validation.Required, validation.Length(6,50)),
	)
}

func (u *User) HasAdminPermission() bool {
	return u.Access > 0
}

func (u *User) Sanitize()  {
	u.Password = ""
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func (u *User) EncryptPassword() (string, error)  {
	b, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
