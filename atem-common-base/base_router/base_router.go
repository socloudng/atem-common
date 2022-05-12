package base_router

import "github.com/gin-gonic/gin"

type ISubRoute interface {
	InitRouter(Router *gin.RouterGroup)
}
