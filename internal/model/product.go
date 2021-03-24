package model

import validation "github.com/go-ozzo/ozzo-validation"

type Product struct {
	ID          int64    `json:"id"`
	Title       string   `json:"title"`
	Weight      float32  `json:"weight"`
	Size        float32  `json:"size"`
	Description string   `json:"description"`
	PhotoLink   string   `json:"photo_link"`
	Price       float32  `json:"price"`
	Status      int8     `json:"status"`
	Category    Category `json:"category,omitempty"`
}

func (p *Product) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.Title, validation.Required),
		validation.Field(&p.Weight, validation.Required),
		validation.Field(&p.Size, validation.Required),
		validation.Field(&p.Description, validation.Required),
		validation.Field(&p.PhotoLink, validation.Required),
		validation.Field(&p.Price, validation.Required),
	)
}
