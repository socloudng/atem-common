package ginhelp

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func SimpleCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowedOrigins: []string{"*"}, // []string{"http://localhost:8080"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Origin", "Range", "x-requested-with", "content-Type"},
		ExposedHeaders: []string{"Content-Length", "Accept-Ranges", "Content-Range", "Content-Disposition"},
		// AllowCredentials: true,
	})
}
