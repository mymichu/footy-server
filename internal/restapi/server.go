package restapi

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func (r* RestAPISettings) Listen(port int) {
	s := gin.Default()
	s.GET("/user/:id", getSingleUser)
	s.GET("/game/round", r.getGameRound)
	s.POST("/game/bet",r.setGameBet)
	
	var portString = ":"+strconv.Itoa(port)
	s.Run(portString) 
}