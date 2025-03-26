package model

type Pulse struct {
	Tenant     string  `json:"tenant" binding:"required"`
	ProductSKU string  `json:"product_sku" binding:"required"`
	UsedAmount float64 `json:"used_amount" binding:"required"`
	UseUnit    string  `json:"use_unit" binding:"required"`
}
