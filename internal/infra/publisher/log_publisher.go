package publisher

import (
	"ingestor/internal/service"

	"go.uber.org/zap"
)

type LogPublisher struct {
	logger *zap.SugaredLogger
}

func NewLogPublisher(logger *zap.SugaredLogger) *LogPublisher {
	return &LogPublisher{
		logger: logger,
	}
}

func (p *LogPublisher) Publish(pulses []service.AggregatedPulseDTO) error {
	p.logger.Infow("Simulated publish of aggregated pulses", "payload", pulses)

	return nil
}
