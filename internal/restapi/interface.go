package restapi

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

type Logic interface {
    CurrentRound() []Game
}

type RestAPISettings struct {
	LogicModule Logic
}