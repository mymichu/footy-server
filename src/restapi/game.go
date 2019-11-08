package restapi

import (
	"github.com/gin-gonic/gin"
)

type Game struct {
	Id string
	HomeTeam string
	GuestTeam string
}

type GameBet struct {
	Id string
	HomeTeamResult string
	GuestTeamResult string
	Joker bool
}

type GameBetResult struct {
	Bet GameBet
	Points struct {
		PickCorrectTeam bool
		PickCorrectDiff bool
		TotalPoints     uint8
	}
}

type Round struct {
    Games []Game 
}


func getGameRound(c *gin.Context) {
	var tmp Game
	tmp.GuestTeam="Warriors"
	tmp.HomeTeam="Sharks"
	tmp.Id="2019R10G1"
	round := Round{make([]Game, 20)}
	for i := 0;  i<20; i++ {
		round.Games[i]=tmp
	}
	c.JSON(200, round)
}
