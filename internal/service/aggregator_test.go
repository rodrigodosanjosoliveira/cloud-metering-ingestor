package service

import (
	"ingestor/internal/core/ports/mocks"
	"ingestor/internal/model"

	"testing"

	"github.com/stretchr/testify/mock"
)

func TestAggregatorService_AddPulse(t *testing.T) {
	t.Run("should aggregate pulses correctly", func(t *testing.T) {
		mockPub := mocks.NewPublisher(t)

		agg := NewAggregatorService(mockPub)

		pulse1 := model.Pulse{
			Tenant:     "tenant1",
			ProductSKU: "sku1",
			UseUnit:    "unit1",
			UsedAmount: 10.0,
		}
		pulse2 := model.Pulse{
			Tenant:     "tenant1",
			ProductSKU: "sku1",
			UseUnit:    "unit1",
			UsedAmount: 5.0,
		}

		agg.AddPulse(pulse1)
		agg.AddPulse(pulse2)

		if len(agg.aggregates) != 1 {
			t.Errorf("expected 1 aggregate, got %d", len(agg.aggregates))
		}

		key := AggregateKey{
			Tenant:     "tenant1",
			ProductSKU: "sku1",
			UseUnit:    "unit1",
		}
		if agg.aggregates[key] != 15.0 {
			t.Errorf("expected aggregate to be 15.0, got %f", agg.aggregates[key])
		}
	})

	t.Run("should handle multiple aggregates", func(t *testing.T) {
		mockPub := mocks.NewPublisher(t)

		agg := NewAggregatorService(mockPub)

		pulse1 := model.Pulse{
			Tenant:     "tenant1",
			ProductSKU: "sku1",
			UseUnit:    "unit1",
			UsedAmount: 10.0,
		}
		pulse2 := model.Pulse{
			Tenant:     "tenant2",
			ProductSKU: "sku2",
			UseUnit:    "unit2",
			UsedAmount: 5.0,
		}

		agg.AddPulse(pulse1)
		agg.AddPulse(pulse2)

		if len(agg.aggregates) != 2 {
			t.Errorf("expected 2 aggregates, got %d", len(agg.aggregates))
		}

		key1 := AggregateKey{
			Tenant:     "tenant1",
			ProductSKU: "sku1",
			UseUnit:    "unit1",
		}
		if agg.aggregates[key1] != 10.0 {
			t.Errorf("expected aggregate to be 10.0, got %f", agg.aggregates[key1])
		}

		key2 := AggregateKey{
			Tenant:     "tenant2",
			ProductSKU: "sku2",
			UseUnit:    "unit2",
		}
		if agg.aggregates[key2] != 5.0 {
			t.Errorf("expected aggregate to be 5.0, got %f", agg.aggregates[key2])
		}
	})
}

func TestAggregatorService_FlushAggregates(t *testing.T) {
	t.Run("should flush aggregates and call publisher", func(t *testing.T) {
		mockPub := mocks.NewPublisher(t)
		mockPub.EXPECT().Publish(mock.Anything).Return(nil)

		agg := NewAggregatorService(mockPub)

		pulse1 := model.Pulse{
			Tenant:     "tenant1",
			ProductSKU: "sku1",
			UseUnit:    "unit1",
			UsedAmount: 10.0,
		}
		pulse2 := model.Pulse{
			Tenant:     "tenant1",
			ProductSKU: "sku1",
			UseUnit:    "unit1",
			UsedAmount: 5.0,
		}

		agg.AddPulse(pulse1)
		agg.AddPulse(pulse2)

		if len(agg.aggregates) != 1 {
			t.Errorf("expected 1 aggregate before flush, got %d", len(agg.aggregates))
		}

		agg.FlushAggregates()

		if len(agg.aggregates) != 0 {
			t.Errorf("expected 0 aggregates after flush, got %d", len(agg.aggregates))
		}

		mockPub.AssertExpectations(t)
	})

	t.Run("should not call publisher if no aggregates", func(t *testing.T) {
		mockPub := mocks.NewPublisher(t)

		agg := NewAggregatorService(mockPub)

		if len(agg.aggregates) != 0 {
			t.Errorf("expected 0 aggregates before flush, got %d", len(agg.aggregates))
		}

		agg.FlushAggregates()

		if len(agg.aggregates) != 0 {
			t.Errorf("expected 0 aggregates after flush, got %d", len(agg.aggregates))
		}

		mockPub.AssertExpectations(t)
	})

	t.Run("should not panic if publisher is nil", func(t *testing.T) {
		agg := NewAggregatorService(nil)

		pulse1 := model.Pulse{
			Tenant:     "tenant1",
			ProductSKU: "sku1",
			UseUnit:    "unit1",
			UsedAmount: 10.0,
		}

		agg.AddPulse(pulse1)

		if len(agg.aggregates) != 1 {
			t.Errorf("expected 1 aggregate before flush, got %d", len(agg.aggregates))
		}

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("expected no panic, but got %v", r)
			}
		}()

		agg.FlushAggregates()

		if len(agg.aggregates) != 0 {
			t.Errorf("expected 0 aggregates after flush, got %d", len(agg.aggregates))
		}
	})
}

func TestAggregatorService_GetAggregatedData(t *testing.T) {
	t.Run("should return aggregated data", func(t *testing.T) {
		mockPub := mocks.NewPublisher(t)

		agg := NewAggregatorService(mockPub)

		pulse1 := model.Pulse{
			Tenant:     "tenant1",
			ProductSKU: "sku1",
			UseUnit:    "unit1",
			UsedAmount: 10.0,
		}
		pulse2 := model.Pulse{
			Tenant:     "tenant1",
			ProductSKU: "sku1",
			UseUnit:    "unit1",
			UsedAmount: 5.0,
		}

		agg.AddPulse(pulse1)
		agg.AddPulse(pulse2)

		data := agg.GetAggregatedData()

		if len(data) != 1 {
			t.Errorf("expected 1 aggregated data, got %d", len(data))
		}

		if data[0].TotalUsed != 15.0 {
			t.Errorf("expected TotalUsed to be 15.0, got %f", data[0].TotalUsed)
		}
	})

	t.Run("should return empty data if no aggregates", func(t *testing.T) {
		mockPub := mocks.NewPublisher(t)

		agg := NewAggregatorService(mockPub)

		data := agg.GetAggregatedData()

		if len(data) != 0 {
			t.Errorf("expected 0 aggregated data, got %d", len(data))
		}
	})

	t.Run("should return empty data if aggregates are cleared", func(t *testing.T) {
		mockPub := mocks.NewPublisher(t)

		mockPub.EXPECT().Publish(mock.Anything).Return(nil)

		agg := NewAggregatorService(mockPub)

		pulse1 := model.Pulse{
			Tenant:     "tenant1",
			ProductSKU: "sku1",
			UseUnit:    "unit1",
			UsedAmount: 10.0,
		}
		pulse2 := model.Pulse{
			Tenant:     "tenant2",
			ProductSKU: "sku2",
			UseUnit:    "unit2",
			UsedAmount: 5.0,
		}

		agg.AddPulse(pulse1)
		agg.AddPulse(pulse2)

		dataBeforeFlush := agg.GetAggregatedData()

		if len(dataBeforeFlush) != 2 {
			t.Errorf("expected 2 aggregated data, got %d", len(dataBeforeFlush))
		}

		agg.FlushAggregates()

		dataAfterFlush := agg.GetAggregatedData()

		if len(dataAfterFlush) != 0 {
			t.Errorf("expected 0 aggregated data after flush, got %d", len(dataAfterFlush))
		}
	})
}
