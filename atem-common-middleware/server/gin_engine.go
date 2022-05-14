package server

import (
	"github.com/gin-gonic/gin"
	"github.com/socloudng/atem-common/configs"
	"go.uber.org/zap"
)

const (
	DefaultEnvironmentDev     = "dev"
	DefaultEnvironmentTest    = "test"
	DefaultEnvironmentRelease = "prod"
)

func NewGinEngine(cfg *configs.ServerConfig, logger *zap.Logger) *gin.Engine {
	initGinMode(cfg.Env)
	r := gin.Default()
	r.GET("/", showCopyRight)
	return r
}

func initGinMode(profile string) {
	gin.SetMode(gin.ReleaseMode)
	// 默认生产
	if profile != "" {
		switch profile {
		case DefaultEnvironmentDev:
			gin.SetMode(gin.DebugMode)
		case DefaultEnvironmentTest:
			gin.SetMode(gin.TestMode)
		case DefaultEnvironmentRelease:
			gin.SetMode(gin.ReleaseMode)
		default:
			gin.SetMode(gin.ReleaseMode)
		}
	}
}

func showCopyRight(c *gin.Context) {
	cr := "Copyright © 2021-2025 By songzb. All rights reserved."
	c.Writer.WriteString(cr)
	c.Done()
}
