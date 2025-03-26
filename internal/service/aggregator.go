package service

import (
	"ingestor/internal/model"
	"sync"
)

type (
	AggregateKey struct {
		Tenant     string
		ProductSKU string
		UseUnit    string
	}

	AggregatorService struct {
		mu         sync.Mutex
		aggregates map[AggregateKey]float64
		publisher  Publisher
	}

	AggregatedPulseDTO struct {
		Tenant     string  `json:"tenant"`
		ProductSKU string  `json:"product_sku"`
		UseUnit    string  `json:"use_unit"`
		TotalUsed  float64 `json:"total_used"`
	}
)

func NewAggregatorService(pub Publisher) *AggregatorService {
	return &AggregatorService{
		aggregates: make(map[AggregateKey]float64),
		publisher:  pub,
	}
}

func (a *AggregatorService) AddPulse(p model.Pulse) {
	key := AggregateKey{
		Tenant:     p.Tenant,
		ProductSKU: p.ProductSKU,
		UseUnit:    p.UseUnit,
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	a.aggregates[key] += p.UsedAmount
}

func (a *AggregatorService) Flush() {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.publisher != nil {
		data := a.flushAndGetCopy()
		_ = a.publisher.Publish(data)
	}

	a.aggregates = make(map[AggregateKey]float64)
}

func (a *AggregatorService) GetAggregatesDTO() []AggregatedPulseDTO {
	a.mu.Lock()
	defer a.mu.Unlock()

	result := make([]AggregatedPulseDTO, 0, len(a.aggregates))
	for k, v := range a.aggregates {
		result = append(result, AggregatedPulseDTO{
			Tenant:     k.Tenant,
			ProductSKU: k.ProductSKU,
			UseUnit:    k.UseUnit,
			TotalUsed:  v,
		})
	}

	return result
}

func (a *AggregatorService) flushAndGetCopy() []AggregatedPulseDTO {
	result := make([]AggregatedPulseDTO, 0, len(a.aggregates))
	for k, v := range a.aggregates {
		result = append(result, AggregatedPulseDTO{
			Tenant:     k.Tenant,
			ProductSKU: k.ProductSKU,
			UseUnit:    k.UseUnit,
			TotalUsed:  v,
		})
	}

	return result
}
