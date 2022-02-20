package util

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, code int,
	errors interface{}, message string, data map[string]interface{}) {
	c.JSON(code, gin.H{
		"code":    code,
		"success": true,
		"error":   errors,
		"message": message,
		"data":    data,
	})
}

func Fail(c *gin.Context, code int,
	errors string, message string, data map[string]interface{}) {
	c.JSON(code, gin.H{
		"code":    code,
		"success": false,
		"error":   errors,
		"message": message,
		"data":    data,
	})
}
