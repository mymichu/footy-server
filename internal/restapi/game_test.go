package restapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"bytes"
	"encoding/json"
	"errors"
)

type Mock struct {
}

func (m Mock) CurrentRound() []Game {
	var tmp Game
	tmp.GuestTeam = "Warriors"
	tmp.HomeTeam = "Sharks"
	tmp.Id = "2019R10G1"
	round := make([]Game, 1)
	for i := 0; i < 1; i++ {
		round[i] = tmp
	}
	return round
}

func (m Mock) InsertBet(bet GameBet) error {
	var err error
	err = nil
	
	if bet.HomeTeamResult > 100 {
		err = errors.New("Home-Team result is over 100")
	}
	
	return err
}

func createRESTTestContext() (*httptest.ResponseRecorder,*gin.Context,*gin.Engine) {
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(func(c *gin.Context) {
		c.Set("profile", "myfakeprofile")
	})
	return resp,c,r
}

func requestBet(bet GameBet) (string, int) {
	resp, c, r := createRESTTestContext()
	const url = "/game/bet"
	rest := RestAPISettings{Mock{}}
	r.POST(url, rest.setGameBet)

	reqString, err := json.Marshal(bet)
	if err == nil {
		c.Request, _ = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqString))
	}
	
	r.ServeHTTP(resp, c.Request)
	return resp.Body.String(), resp.Code
}

func TestGetGameRound(t *testing.T) {
	resp, c, r:= createRESTTestContext()
	const url = "/game/round"
	rest := RestAPISettings{Mock{}}
	r.GET(url, rest.getGameRound)

	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	r.ServeHTTP(resp, c.Request)

	assert.Equal(t, "[{\"Id\":\"2019R10G1\",\"HomeTeam\":\"Sharks\",\"GuestTeam\":\"Warriors\"}]",resp.Body.String())
}

func TestRestAPISettingsSetGameBetGood(t *testing.T) {
	var bet GameBet
	bet.Id = "TestId"
	bet.HomeTeamResult = 80
	bet.GuestTeamResult = 40
	bet.Joker = false

	body, code := requestBet(bet)
	assert.Equal(t, 200, code)
	assert.Equal(t, "{\"Id\":\"TestId\",\"HomeTeamResult\":80,\"GuestTeamResult\":40,\"Joker\":false}",body)
}

func TestRestAPISettingsSetGameBetBad(t *testing.T) {
	var bet GameBet

	bet.Id = "TestId"
	bet.HomeTeamResult = 120
	bet.GuestTeamResult = 40
	bet.Joker = false

	body, code := requestBet(bet)
	assert.Equal(t, 400, code)
	assert.Equal(t, "{\"Error\":\"Home-Team result is over 100\"}", body)
}