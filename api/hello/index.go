package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*Servlet) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
	return
}
