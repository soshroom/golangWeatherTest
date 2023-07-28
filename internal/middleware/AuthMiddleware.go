package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AuthMiddleware(ctx *gin.Context, logger *zap.Logger) {

	user := ctx.GetHeader("User-Access")
	if user == "superuser" {
		ctx.Next()
	} else {
		logger.Info("Unauthorized user")
		ctx.String(401, "{\"message\": \"Not Authorized\"}")
		ctx.Abort()
	}
}
