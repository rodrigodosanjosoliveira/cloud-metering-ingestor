package validator

import (
	"fmt"
	"ingestor/internal/model"

	"github.com/go-playground/validator/v10"
)

type PulseValidator struct {
	validate *validator.Validate
}

func NewPulseValidator() *PulseValidator {
	return &PulseValidator{validate: validator.New()}
}

func (v *PulseValidator) Validate(p *model.Pulse) error {
	if err := v.validate.Struct(p); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return nil
}
