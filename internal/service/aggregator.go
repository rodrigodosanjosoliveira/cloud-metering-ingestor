package service

import (
	"ingestor/internal/core/dto"
	"ingestor/internal/core/ports"
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
		publisher  ports.Publisher
	}
)

func NewAggregatorService(pub ports.Publisher) *AggregatorService {
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

func (a *AggregatorService) FlushAggregates() {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.publisher != nil {
		data := a.flushAndGetCopy()
		if len(data) != 0 {
			_ = a.publisher.Publish(data)
		}
	}

	a.aggregates = make(map[AggregateKey]float64)
}

func (a *AggregatorService) GetAggregatedData() []dto.AggregatedPulse {
	a.mu.Lock()
	defer a.mu.Unlock()

	result := make([]dto.AggregatedPulse, 0, len(a.aggregates))
	for k, v := range a.aggregates {
		result = append(result, dto.AggregatedPulse{
			Tenant:     k.Tenant,
			ProductSKU: k.ProductSKU,
			UseUnit:    k.UseUnit,
			TotalUsed:  v,
		})
	}

	return result
}

func (a *AggregatorService) flushAndGetCopy() []dto.AggregatedPulse {
	result := make([]dto.AggregatedPulse, 0, len(a.aggregates))
	for k, v := range a.aggregates {
		result = append(result, dto.AggregatedPulse{
			Tenant:     k.Tenant,
			ProductSKU: k.ProductSKU,
			UseUnit:    k.UseUnit,
			TotalUsed:  v,
		})
	}

	return result
}
