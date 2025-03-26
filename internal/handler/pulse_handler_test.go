package handler

import (
	"bytes"
	"encoding/json"
	"ingestor/internal/core/dto"
	"ingestor/internal/handler/mocks"
	"ingestor/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"ingestor/internal/infra/logger"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPulseHandler_CreatePulse(t *testing.T) {
	t.Run("should create a pulse and return 201", func(t *testing.T) {
		log := logger.NewLogger()

		mockAggregator := mocks.NewAggregator(t)

		mockAggregator.EXPECT().AddPulse(mock.Anything).Return()

		h := NewPulseHandler(log, mockAggregator)

		router := gin.Default()
		router.POST("/pulses", h.CreatePulse)

		payload := map[string]interface{}{
			"tenant":      "tenant_a",
			"product_sku": "sku123",
			"used_amount": 123.45,
			"use_unit":    "GB",
		}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest("POST", "/pulses", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)
	})

	t.Run("should return a bad request when the payload is invalid", func(t *testing.T) {
		log := logger.NewLogger()

		mockAggregator := mocks.NewAggregator(t)

		h := NewPulseHandler(log, mockAggregator)

		router := gin.Default()
		router.POST("/pulses", h.CreatePulse)

		payload := map[string]interface{}{
			"tenant":      "",
			"product_sku": "sku123",
		}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest("POST", "/pulses", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})
}

func TestPulseHandler_GetAggregates(t *testing.T) {
	t.Run("should return aggregated data", func(t *testing.T) {
		log := logger.NewLogger()

		mockAggregator := mocks.NewAggregator(t)

		mockAggregator.EXPECT().GetAggregatedData().Return([]dto.AggregatedPulse{
			{
				Tenant:     "tenant_a",
				ProductSKU: "sku123",
				UseUnit:    "GB",
				TotalUsed:  10,
			},
		})

		h := NewPulseHandler(log, mockAggregator)

		router := gin.Default()
		router.GET("/aggregates", h.GetAggregates)

		req, _ := http.NewRequest("GET", "/aggregates", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Contains(t, resp.Body.String(), "tenant_a")
		assert.Contains(t, resp.Body.String(), "sku123")
	})

	t.Run("should return empty data when no aggregates are present", func(t *testing.T) {
		log := logger.NewLogger()

		mockAggregator := mocks.NewAggregator(t)

		mockAggregator.EXPECT().GetAggregatedData().Return([]dto.AggregatedPulse{})

		h := NewPulseHandler(log, mockAggregator)

		router := gin.Default()
		router.GET("/aggregates", h.GetAggregates)

		req, _ := http.NewRequest("GET", "/aggregates", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, "[]", resp.Body.String())
	})

	t.Run("should return 500 on aggregator error", func(t *testing.T) {
		log := logger.NewLogger()

		mockAggregator := mocks.NewAggregator(t)

		mockAggregator.EXPECT().GetAggregatedData().Return(nil)

		h := NewPulseHandler(log, mockAggregator)

		router := gin.Default()
		router.GET("/aggregates", h.GetAggregates)

		req, _ := http.NewRequest("GET", "/aggregates", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestPulseHandler_FlushAggregates(t *testing.T) {
	t.Run("should flush aggregates and return 200", func(t *testing.T) {
		log := logger.NewLogger()

		mockAggregator := mocks.NewAggregator(t)

		mockAggregator.EXPECT().FlushAggregates().Return()

		h := NewPulseHandler(log, mockAggregator)

		router := gin.Default()
		router.POST("/flush", h.FlushAggregates)

		req, _ := http.NewRequest("POST", "/flush", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("should return 500 on flush error", func(t *testing.T) {
		log := logger.NewLogger()

		mockAggregator := mocks.NewAggregator(t)

		mockAggregator.EXPECT().FlushAggregates().Return()

		h := NewPulseHandler(log, mockAggregator)

		router := gin.Default()
		router.POST("/flush", h.FlushAggregates)

		req, _ := http.NewRequest("POST", "/flush", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})
}

func TestCreatePulse_Success(t *testing.T) {
	router := gin.Default()

	log := logger.NewLogger()

	mockAggregator := mocks.NewAggregator(t)

	mockAggregator.EXPECT().AddPulse(mock.Anything).Return()

	h := NewPulseHandler(log, mockAggregator)

	router.POST("/pulses", h.CreatePulse)

	payload := map[string]interface{}{
		"tenant":      "tenant_a",
		"product_sku": "sku123",
		"used_amount": 123.45,
		"use_unit":    "GB",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/pulses", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestCreatePulse_InvalidPayload(t *testing.T) {
	log := logger.NewLogger()

	mockAggregator := mocks.NewAggregator(t)

	h := NewPulseHandler(log, mockAggregator)

	router := gin.Default()
	router.POST("/pulses", h.CreatePulse)

	payload := map[string]interface{}{
		"tenant":      "",
		"product_sku": "sku123",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/pulses", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestGetAggregates(t *testing.T) {
	log := logger.NewLogger()

	mockAggregator := mocks.NewAggregator(t)

	mockAggregator.EXPECT().AddPulse(model.Pulse{
		Tenant:     "tenant_a",
		ProductSKU: "sku123",
		UsedAmount: 10,
		UseUnit:    "GB",
	}).Return()

	mockAggregator.EXPECT().GetAggregatedData().Return([]dto.AggregatedPulse{
		{
			Tenant:     "tenant_a",
			ProductSKU: "sku123",
			UseUnit:    "GB",
			TotalUsed:  10,
		},
	})

	h := NewPulseHandler(log, mockAggregator)

	router := gin.Default()
	router.POST("/flush", h.FlushAggregates)
	router.POST("/pulses", h.CreatePulse)
	router.GET("/aggregates", h.GetAggregates)

	payload := map[string]interface{}{
		"tenant":      "tenant_a",
		"product_sku": "sku123",
		"used_amount": 10,
		"use_unit":    "GB",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/pulses", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(httptest.NewRecorder(), req)

	reqGet, _ := http.NewRequest("GET", "/aggregates", nil)
	respGet := httptest.NewRecorder()
	router.ServeHTTP(respGet, reqGet)

	assert.Equal(t, http.StatusOK, respGet.Code)
	assert.Contains(t, respGet.Body.String(), "tenant_a")
	assert.Contains(t, respGet.Body.String(), "sku123")
}

func TestFlushAggregates(t *testing.T) {
	log := logger.NewLogger()

	mockAggregator := mocks.NewAggregator(t)

	mockAggregator.EXPECT().AddPulse(model.Pulse{
		Tenant:     "tenant_flush_test",
		ProductSKU: "sku_flush",
		UsedAmount: 5,
		UseUnit:    "GB",
	}).Return()

	mockAggregator.EXPECT().FlushAggregates().Return().Once()

	mockAggregator.EXPECT().GetAggregatedData().Return([]dto.AggregatedPulse{
		{
			Tenant:     "tenant_flush_test",
			ProductSKU: "sku_flush",
			UseUnit:    "GB",
			TotalUsed:  5,
		},
	})

	h := NewPulseHandler(log, mockAggregator)

	router := gin.Default()
	router.POST("/flush", h.FlushAggregates)
	router.POST("/pulses", h.CreatePulse)
	router.GET("/aggregates", h.GetAggregates)

	payload := map[string]interface{}{
		"tenant":      "tenant_flush_test",
		"product_sku": "sku_flush",
		"used_amount": 5,
		"use_unit":    "GB",
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/pulses", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(httptest.NewRecorder(), req)

	reqGetBefore, _ := http.NewRequest("GET", "/aggregates", nil)
	respBefore := httptest.NewRecorder()
	router.ServeHTTP(respBefore, reqGetBefore)
	assert.Contains(t, respBefore.Body.String(), "tenant_flush_test")

	reqFlush, _ := http.NewRequest("POST", "/flush", nil)
	respFlush := httptest.NewRecorder()
	router.ServeHTTP(respFlush, reqFlush)
	assert.Equal(t, http.StatusOK, respFlush.Code)
}

func TestFlush_CallsFlushOnAggregator(t *testing.T) {
	log := logger.NewLogger()

	mockAggregator := mocks.NewAggregator(t)

	mockAggregator.EXPECT().FlushAggregates().Return().Once()

	h := NewPulseHandler(log, mockAggregator)

	r := gin.Default()
	r.POST("/flush", h.FlushAggregates)

	req, _ := http.NewRequest("POST", "/flush", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
}
