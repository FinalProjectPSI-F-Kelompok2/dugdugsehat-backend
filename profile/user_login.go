package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginProfile struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

type RegisterProfile struct {
	Email    string `json: "email"`
	Username string `json: "username"`
	Password string `json: "password"`
}

func Login() gin.HandlerFunc {
	// Serve to the client
	return func(c *gin.Context) {
		var lp LoginProfile
		if c.BindJSON(&lp) != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error in server",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"valid": true,
		})
	}
}

func Register() {

}
