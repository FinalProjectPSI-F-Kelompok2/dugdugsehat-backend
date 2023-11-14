package profile

import (
	"net/http"

	"github.com/FinalProjectPSI-F-Kelompok2/dugdugsehat-backend/model"
	"github.com/gin-gonic/gin"
)

type UserProfile struct {
	User struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"user"`
	HealthProfile struct {
		GenderIsMale bool `json:"genderIsMale"`
		Age          int  `json:"age"`
		BodyHeight   int  `json:"bodyHeight"`
		BodyWeight   int  `json:"bodyWeight"`
	} `json:"health_profile"`
}

func UpdateProfile(db *model.DbCon) gin.HandlerFunc {
	return func(c *gin.Context) {
		var up UserProfile
		if c.BindJSON(&up) != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error in server",
			})
			return
		}
		r, err := db.Db.Query(
			"UPDATE profile SET full_name=$2, sex=$3, age=$4, body_height=$5, body_weight=$6 WHERE email=$1",
			up.User.Email,
			up.User.Name,
			up.HealthProfile.GenderIsMale,
			up.HealthProfile.Age,
			up.HealthProfile.BodyHeight,
			up.HealthProfile.BodyWeight,
		)
		defer r.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "db error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "edit ok",
		})
	}
}

func GetProfile(db *model.DbCon) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.GetHeader("email")
		r := db.Db.QueryRow("SELECT * FROM profile WHERE email=$1", email)
		var up UserProfile
		r.Scan(
			&up.User.Email,
			&up.User.Name,
			&up.HealthProfile.BodyHeight,
			&up.HealthProfile.BodyWeight,
			&up.HealthProfile.Age,
			&up.HealthProfile.GenderIsMale,
		)
		if len(up.User.Email) <= 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "user not found",
			})
			return
		}
		c.JSON(http.StatusOK, &up)
	}
}
