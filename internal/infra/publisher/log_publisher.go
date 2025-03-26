package publisher

import (
	"ingestor/internal/core/dto"

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

func (p *LogPublisher) Publish(pulses []dto.AggregatedPulse) error {
	p.logger.Infow("Simulated publish of aggregated pulses", "payload", pulses)

	return nil
}
