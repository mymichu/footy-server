package mock

import (
	"github.com/mymichu/footy-server/internal/restapi"
)

type MockLogic struct{
}

func (m MockLogic) CurrentRound() []restapi.Game {
	var tmp restapi.Game
	tmp.GuestTeam="Warriors"
	tmp.HomeTeam="Sharks"
	tmp.Id="2019R10G1"
	round := make([]restapi.Game, 20)
	for i := 0;  i<20; i++ {
		round[i]=tmp
	}
	return round
}