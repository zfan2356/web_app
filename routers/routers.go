package routers

import (
	"github.com/gin-gonic/gin"
	"web_app/logger"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	return r
}
