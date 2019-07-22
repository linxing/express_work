package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(s Servlet) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/index", s.Index)
	}

	return r
}
