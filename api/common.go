package api

import "github.com/gin-gonic/gin"

func handleError(c *gin.Context, error error) {
	if error != nil {
		c.JSON(422, gin.H{"status": "error", "message": error.Error()})
	}
}
