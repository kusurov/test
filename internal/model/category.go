package model

import validation "github.com/go-ozzo/ozzo-validation"

type Category struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Status int8   `json:"status"`
}

func (c *Category) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Title, validation.Required),
	)
}
