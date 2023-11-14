package profile

import (
	"net/http"
	"strings"

	"github.com/FinalProjectPSI-F-Kelompok2/dugdugsehat-backend/model"
	"github.com/gin-gonic/gin"
)

type LoginProfile struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

type RegisterProfile struct {
	Email    string `json: "email"`
	Name     string `json: "name"`
	Password string `json: "password"`
}

func Login(db *model.DbCon) gin.HandlerFunc {
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

func Register(db *model.DbCon) gin.HandlerFunc {
	return func(c *gin.Context) {
		var rp RegisterProfile
		if c.BindJSON(&rp) != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "server error",
			})
			return
		}
		rp.Email = strings.ToLower(rp.Email)

		if len(rp.Email) == 0 || len(rp.Password) == 0 || len(rp.Name) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "some fields are empty",
			})
			return
		}

		// Check if email is not used
		var emails string
		db.Db.QueryRow("SELECT email FROM users WHERE email=$1 LIMIT 1", rp.Email).Scan(&emails)
		if len(emails) > 0 {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "email exists",
			})
			return
		}

		if r, err := db.Db.Query("INSERT INTO users VALUES ($1, $2)", rp.Email, rp.Password); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "db error",
			})
			return
		} else {
			r.Close()
		}

		if r, err := db.Db.Query("INSERT INTO profile VALUES ($1, $2, null, null, false)", rp.Email, rp.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "db error",
			})
			return
		} else {
			r.Close()
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}
