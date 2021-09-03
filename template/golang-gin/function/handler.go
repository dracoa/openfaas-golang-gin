package function

import (
	"github.com/gin-gonic/gin"
	_ "time/tzdata"
)

func Handle(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
