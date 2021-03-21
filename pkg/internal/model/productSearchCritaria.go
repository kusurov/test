package model

type ProductSearchCriteria struct {
	SearchCriteria struct {
		Title string `json:"title"`
	} `json:"searchCriteria"`
}
