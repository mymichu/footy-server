package mock

import (
	"github.com/mymichu/footy-server/internal/restapi"
	"errors"
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

func (m MockLogic) InsertBet(bet restapi.GameBet) error {
	var tmp restapi.GameBet
	var err error
	err = nil
	
	if tmp.HomeTeamResult > 100 {
		err = errors.New("Home-Team result is over 100")
	}
	
	return err
}