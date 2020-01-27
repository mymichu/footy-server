package restapi

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)


func TestGetSingleUser( t *testing.T ){

	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(func(c *gin.Context) {
		c.Set("profile", "myfakeprofile")
	})
	r.GET("/user/:id", getSingleUser)
	
	c.Request, _ = http.NewRequest(http.MethodGet, "/user/101", nil)
	r.ServeHTTP(resp, c.Request)

	assert.Equal(t, resp.Body.String(), "{\"message\":\"User 101\"}")
}
