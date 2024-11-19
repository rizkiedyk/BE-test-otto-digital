package dto

type ReqBrand struct {
	BrandID string `json:"brand_id"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
}
