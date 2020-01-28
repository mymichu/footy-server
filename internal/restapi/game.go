package restapi

import (
	"github.com/gin-gonic/gin"
)

type Round struct {
    Games []Game 
}


func (r* RestAPISettings) getGameRound(c *gin.Context) {
	round := r.LogicModule.CurrentRound();
	c.JSON(200, round)
}

func (r* RestAPISettings) setGameBet(c *gin.Context) {
	var bet GameBet
	var clientError ClientError
	err := c.BindJSON(&bet)
	if err == nil {
		err = r.LogicModule.InsertBet(bet)
		if err == nil {
			c.JSON(200, bet)
		} 
		
	} 

	if err != nil {
		clientError.Error = err.Error()
		c.JSON(400,clientError)
	}
}
