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

func TestGetGameRound(t *testing.T) {
	rest := RestAPISettings{Mock{}}

	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(func(c *gin.Context) {
		c.Set("profile", "myfakeprofile")
	})
	r.GET("/game/round", rest.getGameRound)

	c.Request, _ = http.NewRequest(http.MethodGet, "/game/round", nil)
	r.ServeHTTP(resp, c.Request)

	var ref Game
	ref.GuestTeam = "Warriors"
	ref.HomeTeam = "Sharks"
	ref.Id = "2019R10G1"
	round := make([]Game, 1)
	for i := 0; i < 1; i++ {
		round[i] = ref
	}

	assert.Equal(t, "[{\"Id\":\"2019R10G1\",\"HomeTeam\":\"Sharks\",\"GuestTeam\":\"Warriors\"}]",resp.Body.String())
}


func TestRestAPISettings_setGameBetGood(t *testing.T) {
	rest := RestAPISettings{Mock{}}

	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(func(c *gin.Context) {
		c.Set("profile", "myfakeprofile")
	})
	r.POST("/game/bet", rest.setGameBet)

	var bet GameBet

	bet.Id = "TestId"
	bet.HomeTeamResult = 80
	bet.GuestTeamResult = 40
	bet.Joker = false

	reqString, err := json.Marshal(bet)
	if err == nil {
		c.Request, _ = http.NewRequest(http.MethodPost, "/game/bet", bytes.NewBuffer(reqString))
	}
	
	r.ServeHTTP(resp, c.Request)
	assert.Equal(t, 200, resp.Code)
	body := resp.Body.String()
	assert.Equal(t, "{\"Id\":\"TestId\",\"HomeTeamResult\":80,\"GuestTeamResult\":40,\"Joker\":false}",body)
}

func TestRestAPISettings_setGameBetBad(t *testing.T) {
	rest := RestAPISettings{Mock{}}

	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(func(c *gin.Context) {
		c.Set("profile", "myfakeprofile")
	})
	r.POST("/game/bet", rest.setGameBet)

	var bet GameBet

	bet.Id = "TestId"
	bet.HomeTeamResult = 120
	bet.GuestTeamResult = 40
	bet.Joker = false

	reqString, err := json.Marshal(bet)
	if err == nil {
		c.Request, _ = http.NewRequest(http.MethodPost, "/game/bet", bytes.NewBuffer(reqString))
	}
	
	r.ServeHTTP(resp, c.Request)
	assert.Equal(t, 400, resp.Code)
	body := resp.Body.String()
	assert.Equal(t, "{\"Error\":\"Home-Team result is over 100\"}", body)
}