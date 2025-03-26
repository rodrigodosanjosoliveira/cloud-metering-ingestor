package handler

import (
	"ingestor/internal/infra/publisher"
	"ingestor/internal/infra/validator"
	"ingestor/internal/model"
	"ingestor/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

type PulseHandler struct {
	logger     *zap.SugaredLogger
	aggregator *service.AggregatorService
}

func NewPulseHandler(logger *zap.SugaredLogger) *PulseHandler {
	pub := publisher.NewLogPublisher(logger)
	aggregator := service.NewAggregatorService(pub)

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
		h.logger.Warn("Validation error", zap.Error(err))

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	h.aggregator.AddPulse(pulse)

	c.Status(http.StatusCreated)
}

// GetAggregates godoc
// @Summary Get aggregated usage data
// @Description Returns current aggregation grouped by tenant, SKU and unit
// @Tags pulses
// @Produce json
// @Success 200 {array} service.AggregatedPulseDTO
// @Router /aggregates [get]
func (h *PulseHandler) GetAggregates(c *gin.Context) {
	aggregates := h.aggregator.GetAggregatesDTO()

	c.JSON(http.StatusOK, aggregates)
}

// FlushAggregates godoc
// @Summary Flush current aggregated data
// @Description Simulates sending the aggregated data and clears current state
// @Tags pulses
// @Success 200 {string} string "Flushed"
// @Router /flush [post]
func (h *PulseHandler) FlushAggregates(c *gin.Context) {
	h.aggregator.Flush()

	c.Status(http.StatusOK)
}
