package validator

import (
	"ingestor/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPulseValidator_ValidPulse(t *testing.T) {
	v := NewPulseValidator()

	p := &model.Pulse{
		Tenant:     "tenant123",
		ProductSKU: "sku456",
		UseUnit:    "GB",
		UsedAmount: 120,
	}

	err := v.Validate(p)
	assert.NoError(t, err)
}

func TestPulseValidator_InvalidPulse(t *testing.T) {
	v := NewPulseValidator()

	p := &model.Pulse{
		Tenant:     "",
		ProductSKU: "",
		UsedAmount: 0,
		UseUnit:    "",
	}

	err := v.Validate(p)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed")
}

func TestPulseValidator_InvalidUseUnit(t *testing.T) {
	v := NewPulseValidator()

	p := &model.Pulse{
		Tenant:     "tenant123",
		ProductSKU: "sku456",
		UsedAmount: 120,
		UseUnit:    "INVALID_UNIT",
	}

	err := v.Validate(p)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed")
}

func TestPulseValidator_ZeroUsedAmount(t *testing.T) {
	v := NewPulseValidator()

	p := &model.Pulse{
		Tenant:     "tenant123",
		ProductSKU: "sku456",
		UsedAmount: 0,
		UseUnit:    "GB",
	}

	err := v.Validate(p)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed")
}
