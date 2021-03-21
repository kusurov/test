package model

type Catalog struct {
	Products []*Product `json:"products"`
}
