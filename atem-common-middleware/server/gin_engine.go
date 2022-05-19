package server

import (
	"github.com/gin-gonic/gin"
)

const (
	DefaultEnvironmentDev     = "dev"
	DefaultEnvironmentTest    = "test"
	DefaultEnvironmentRelease = "prod"
)

func NewGinEngine(profile string) *gin.Engine {
	initGinMode(profile)
	r := gin.Default()
	r.GET("/copyRight", func(c *gin.Context) {
		cr := "Copyright © 2021-2025 By songzb. All rights reserved."
		c.Writer.WriteString(cr)
		c.Done()
	})
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
