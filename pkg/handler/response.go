package handler

import "github.com/gin-gonic/gin"

func responseWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"status":      "error",
		"description": message,
	})
}

func responseSuccessful(c *gin.Context, message interface{}) {
	c.AbortWithStatusJSON(200, gin.H{
		"status":      "success",
		"description": message,
	})
}
