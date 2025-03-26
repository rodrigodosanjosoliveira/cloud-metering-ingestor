package dto

type AggregatedPulse struct {
	Tenant     string  `json:"tenant"`
	ProductSKU string  `json:"product_sku"`
	UseUnit    string  `json:"use_unit"`
	TotalUsed  float64 `json:"total_used"`
}
