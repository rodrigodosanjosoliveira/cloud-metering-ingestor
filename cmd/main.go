package main

import (
	"ingestor/internal/handler"
	"ingestor/internal/infra/logger"
	"ingestor/internal/infra/metrics"
	"ingestor/internal/infra/publisher"
	"ingestor/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"

	_ "ingestor/docs" // docs are generated by Swag CLI, you have to import it.

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Cloud Metering Ingestor API
// @version 1.0
// @description API for usage pulse ingestion and aggregation.
// @BasePath /
func main() {
	log := logger.NewLogger()
	defer func(log *zap.SugaredLogger) {
		err := log.Sync()
		if err != nil {
			log.Error("Failed to sync logger", err)
		}
	}(log)

	log.Info("Starting the application...")

	r := gin.Default()

	metrics.Init()

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	pub := publisher.NewLogPublisher(log)
	aggregator := service.NewAggregatorService(pub)

	h := handler.NewPulseHandler(log, aggregator)

	r.POST("/pulses", h.CreatePulse)
	r.GET("/aggregates", h.GetAggregates)
	r.POST("/flush", h.FlushAggregates)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Info("Server is running on port 8080")

	if err := r.Run(":8080"); err != nil {
		log.Error("Failed to start server", err)
	}

	log.Info("Server stopped")
}
