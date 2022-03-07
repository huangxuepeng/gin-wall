package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	ids := c.Param("id")
	c.JSON(200, gin.H{
		"id": ids,
	})
	fmt.Println(ids)
}
