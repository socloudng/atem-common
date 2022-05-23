package ginhelp

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func MatchPrefixFileHandler(prefixes []string, filepath string, httpMethods ...string) gin.HandlerFunc {
	httpMethodStrs := strings.Join(httpMethods, ",")
	return func(c *gin.Context) {
		for _, prefix := range prefixes {
			if strings.HasPrefix(c.Request.RequestURI, prefix) &&
				(httpMethodStrs == "" || strings.Contains(httpMethodStrs, c.Request.Method)) {
				c.File(filepath)
				break
			} else {
				c.Next()
			}
		}
	}
}

func MatchPrefixDoHandler(prefixes []string, matchFunc gin.HandlerFunc, httpMethods ...string) gin.HandlerFunc {
	httpMethodStrs := strings.Join(httpMethods, ",")
	return func(c *gin.Context) {
		for _, prefix := range prefixes {
			if strings.HasPrefix(c.Request.RequestURI, prefix) &&
				(httpMethodStrs == "" || strings.Contains(httpMethodStrs, c.Request.Method)) {
				matchFunc(c)
				break
			} else {
				c.Next()
			}
		}
	}
}
