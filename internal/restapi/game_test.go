package restapi

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

type Mock struct{
}

func (m Mock) CurrentRound() []Game {
	var tmp Game
	tmp.GuestTeam="Warriors"
	tmp.HomeTeam="Sharks"
	tmp.Id="2019R10G1"
	round := make([]Game, 1)
	for i := 0;  i<1; i++ {
		round[i]=tmp
	}
	return round
}

func TestGetGameRound( t *testing.T ){
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
	ref.GuestTeam="Warriors"
	ref.HomeTeam="Sharks"
	ref.Id="2019R10G1"
	round := make([]Game, 1)
	for i := 0;  i<1; i++ {
		round[i]=ref
	}

	assert.Equal(t, resp.Body.String(), "[{\"Id\":\"2019R10G1\",\"HomeTeam\":\"Sharks\",\"GuestTeam\":\"Warriors\"}]")
}
