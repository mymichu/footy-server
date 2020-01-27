package restapi

import (
	"github.com/gin-gonic/gin"
)


func getSingleUser(c *gin.Context) {
	id := c.Param("id")
    user := "User "+id 
	c.JSON(200, gin.H{
		"message": user,
	})
}
