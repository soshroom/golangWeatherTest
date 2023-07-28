package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golangWeatherTest/internal/domain"
	"golangWeatherTest/internal/middleware"
)

type mainController struct {
	mathService IMathService
	config      *domain.Config
	logger      *zap.Logger
	engine      *gin.Engine
}

func NewMainController(mathService IMathService, logger *zap.Logger, config *domain.Config) *mainController {
	controller := &mainController{
		mathService: mathService,
		config:      config,
		logger:      logger,
		engine:      gin.Default(),
	}

	controller.initRoutes()

	controller.engine.Run(config.Address)

	return controller
}

func (c mainController) mathRequest(ctx *gin.Context) {
	expression := ctx.Query("expression")
	result, err := c.mathService.ProcessExpression(expression)
	if err != nil || expression == "" {
		ctx.String(400, "{\"message\": \"bad expression\"}")
		return
	}
	res, _ := json.Marshal(struct {
		Result int `json:"result"`
	}{
		Result: result,
	})
	ctx.String(200, string(res))
}

func (c mainController) initRoutes() {
	userAccessGroup := c.engine.Group("", func(ctx *gin.Context) {
		middleware.AuthMiddleware(ctx, c.logger)
	})
	userAccessGroup.GET("/math", c.mathRequest)
}
