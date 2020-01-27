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
