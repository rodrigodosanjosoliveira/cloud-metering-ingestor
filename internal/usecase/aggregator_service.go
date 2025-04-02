package usecase

import (
	"ingestor/internal/core/ports"
	"ingestor/internal/handler"
	"log"
	"time"
)

type AggregatorService struct {
	aggregator handler.Aggregator
	publisher  ports.Publisher
}

func NewAggregatorService(aggregator handler.Aggregator, publisher ports.Publisher) *AggregatorService {
	return &AggregatorService{
		aggregator: aggregator,
		publisher:  publisher,
	}
}

func (s *AggregatorService) StartPeriodicFlush(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			<-ticker.C
			s.FlushAggregates()
		}
	}()
}

func (s *AggregatorService) FlushAggregates() {
	log.Println("[flush] Starting periodic flush...")

	pulses := s.aggregator.GetAggregatedData()

	if len(pulses) == 0 {
		log.Println("[flush] No data to send.")

		return
	}

	err := s.publisher.Publish(pulses)
	if err != nil {
		log.Printf("[flush] Error publishing data: %v\n", err)

		return
	}

	log.Printf("[flush] %d aggregated events published successfully.\n", len(pulses))
}
