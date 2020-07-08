package function

import (
	"github.com/gin-gonic/gin"
)

func Handle(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
