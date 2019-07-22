package routers

import (
	"github.com/gin-gonic/gin"
)

type Servlet interface {
	Hello
}

type Hello interface {
	Index(c *gin.Context)
}
