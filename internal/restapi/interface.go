package restapi

type ClientError struct {
	Error string
}
type Game struct {
	Id string
	HomeTeam string
	GuestTeam string
}

type GameBet struct {
	Id string
	HomeTeamResult uint8
	GuestTeamResult uint8
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
	InsertBet(bet GameBet) error
}

type RestAPISettings struct {
	LogicModule Logic
}