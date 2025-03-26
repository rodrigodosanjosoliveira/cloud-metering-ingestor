package service

type Publisher interface {
	Publish(data []AggregatedPulseDTO) error
}
