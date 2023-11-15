package profile

import (
	"net/http"
	"strings"

	"github.com/FinalProjectPSI-F-Kelompok2/dugdugsehat-backend/model"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

type LoginProfile struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterProfile struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
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

		var dblp LoginProfile
		db.Db.QueryRow("SELECT * FROM users WHERE email=$1 LIMIT 1", lp.Email).Scan(&dblp.Email, &dblp.Password)

		var credentialsValidStr string
		if dblp.Password == lp.Password {
			credentialsValidStr = "user valid"
		} else {
			credentialsValidStr = "user invalid"
		}

		c.JSON(http.StatusOK, gin.H{
			"status": credentialsValidStr,
		})
	}
}

func Register(db *model.DbCon) gin.HandlerFunc {
	return func(c *gin.Context) {
		var rp RegisterProfile
		if c.BindJSON(&rp) != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "server error",
			})
			return
		}
		rp.Email = strings.ToLower(rp.Email)

		if len(rp.Email) == 0 || len(rp.Password) == 0 || len(rp.Name) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "some fields are empty",
			})
			return
		}

		// Check if email is not used
		var emails string
		db.Db.QueryRow("SELECT email FROM users WHERE email=$1 LIMIT 1", rp.Email).Scan(&emails)
		if len(emails) > 0 {
			c.JSON(http.StatusForbidden, gin.H{
				"status": "email exists",
			})
			return
		}

		var r *pgx.Rows
		var err error

		r, err = db.Db.Query("INSERT INTO users VALUES ($1, $2)", rp.Email, rp.Password)
		r.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "db error",
			})
			return
		}

		r, err = db.Db.Query("INSERT INTO profile VALUES ($1, $2, null, null, null, false)", rp.Email, rp.Name)
		r.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "db error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}
