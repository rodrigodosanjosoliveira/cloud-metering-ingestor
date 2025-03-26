package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ingestor/internal/infra/logger"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	log := logger.NewLogger()
	h := NewPulseHandler(log)

	r.POST("/pulses", h.CreatePulse)
	r.GET("/aggregates", h.GetAggregates)
	r.POST("/flush", h.FlushAggregates)
	return r
}

func TestCreatePulse_Success(t *testing.T) {
	router := setupRouter()

	payload := map[string]interface{}{
		"tenant":      "tenant_a",
		"product_sku": "sku123",
		"used_amount": 123.45,
		"use_unit":    "GB",
		"timestamp":   "2025-03-21T10:30:00Z",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/pulses", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestCreatePulse_InvalidPayload(t *testing.T) {
	router := setupRouter()

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
	router := setupRouter()

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
	router := setupRouter()

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

	reqGetAfter, _ := http.NewRequest("GET", "/aggregates", nil)
	respAfter := httptest.NewRecorder()
	router.ServeHTTP(respAfter, reqGetAfter)
	assert.NotContains(t, respAfter.Body.String(), "tenant_flush_test")
}
