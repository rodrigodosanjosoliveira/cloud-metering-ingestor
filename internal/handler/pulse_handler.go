package handler

import (
	"ingestor/internal/core/dto"
	"ingestor/internal/infra/metrics"
	"ingestor/internal/infra/validator"
	"ingestor/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

//go:generate mockery --name=Aggregator --output=./mocks --with-expecter
type Aggregator interface {
	AddPulse(pulse model.Pulse)
	GetAggregatedData() []dto.AggregatedPulse
	FlushAggregates()
}

type PulseHandler struct {
	logger     *zap.SugaredLogger
	aggregator Aggregator
}

func NewPulseHandler(logger *zap.SugaredLogger, aggregator Aggregator) *PulseHandler {
	return &PulseHandler{
		logger:     logger,
		aggregator: aggregator,
	}
}

// CreatePulse godoc
// @Summary Receive a new usage pulse
// @Description Accepts a usage pulse and adds it to the aggregation service
// @Tags pulses
// @Accept json
// @Produce json
// @Param pulse body model.Pulse true "Usage Pulse Data"
// @Success 201 {string} string "Created"
// @Failure 400 {object} map[string]string
// @Router /pulses [post]
func (h *PulseHandler) CreatePulse(c *gin.Context) {
	var pulse model.Pulse

	val := validator.NewPulseValidator()

	if err := c.ShouldBindJSON(&pulse); err != nil {
		h.logger.Warn("Invalid JSON payload.", zap.Error(err))

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := val.Validate(&pulse); err != nil {
		h.logger.Warn("Invalid pulse data.", zap.Error(err))

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	h.aggregator.AddPulse(pulse)

	h.logger.Infow("Pulse added to aggregator", "pulse", pulse)

	//TODO: ver open telemetry
	metrics.PulsesReceived.Inc()

	c.Status(http.StatusCreated)
}

// GetAggregates godoc
// @Summary Get aggregated usage data
// @Description Returns current aggregation grouped by tenant, SKU and unit
// @Tags pulses
// @Produce json
// @Success 200 {array} dto.AggregatedPulse
// @Router /aggregates [get]
func (h *PulseHandler) GetAggregates(c *gin.Context) {
	aggregates := h.aggregator.GetAggregatedData()

	if aggregates == nil {
		h.logger.Error("Failed to get aggregated data")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get aggregated data"})

		return
	}

	metrics.AggregationCount.Set(float64(len(aggregates)))

	h.logger.Infow("Aggregated data retrieved", "aggregates", aggregates)

	c.JSON(http.StatusOK, aggregates)
}

// FlushAggregates godoc
// @Summary Flush current aggregated data
// @Description Simulates sending the aggregated data and clears current state
// @Tags pulses
// @Success 200 {string} string "Flushed"
// @Router /flush [post]
func (h *PulseHandler) FlushAggregates(c *gin.Context) {
	h.aggregator.FlushAggregates()

	h.logger.Infow("Aggregated data flushed")

	metrics.FlushTotal.Inc()

	metrics.AggregationCount.Set(0)

	c.JSON(http.StatusOK, gin.H{"status": "Flushed"})
}
