package model

type Pulse struct {
	Tenant     string  `json:"tenant" binding:"required" validate:"required"`
	ProductSKU string  `json:"product_sku" binding:"required" validate:"required"`
	UsedAmount float64 `json:"used_amount" binding:"required" validate:"required,gt=0"`
	UseUnit    string  `json:"use_unit" binding:"required" validate:"required,oneof=GB MB KB TB PB"`
}
