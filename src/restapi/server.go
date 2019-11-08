package restapi

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func Listen(port int) {
	r := gin.Default()
	r.GET("/user/:id", getSingleUser)
	r.GET("/game/round", getGameRound)
	
	var portString = ":"+strconv.Itoa(port)
	r.Run(portString) 
}