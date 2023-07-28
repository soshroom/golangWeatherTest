package app

import (
	"go.uber.org/zap"
	controller "golangWeatherTest/internal/controller"
	"golangWeatherTest/internal/domain"
	"golangWeatherTest/internal/service"
)

func Init(cfg *domain.Config) {

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	logger.Info("cfg")

	MathService := service.NewMathService()

	_ = controller.NewMainController(MathService, logger, cfg)

}
