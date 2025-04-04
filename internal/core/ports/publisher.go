package ports

import "ingestor/internal/core/dto"

//go:generate mockery --name=Publisher --output=./mocks --with-expecter
type Publisher interface {
	Publish(aggregatedPulse []dto.AggregatedPulse) error
}
