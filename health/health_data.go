package health

import (
	"fmt"
	"net/http"
	"time"

	"github.com/FinalProjectPSI-F-Kelompok2/dugdugsehat-backend/model"
	"github.com/gin-gonic/gin"
)

type SaveDataRequest struct {
	Email string `json:"email"`
	Data  struct {
		Ecg float32 `json:"ecg"`
		Hr  int     `json:"hr"`
	} `json:"data"`
}

type HealthData struct {
	Date time.Time `sql:"measureDate"`
	Ecg  float32   `sql:"ecg"`
	Hr   int       `sql:"heartRate"`
}

func SaveHealthData(db *model.DbCon) gin.HandlerFunc {
	return func(c *gin.Context) {
		var sdr SaveDataRequest
		if c.BindJSON(&sdr) != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error in server",
			})
			return
		}
		dataTimestamp := time.Now()
		r, err := db.Db.Query("INSERT INTO health_data VALUES ($1, $2, $3, $4)", sdr.Email, dataTimestamp, sdr.Data.Ecg, sdr.Data.Hr)
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

func GetHealthData(db *model.DbCon) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.GetHeader("email")
		rowLimit := c.GetHeader("row")

		// check if email is registered
		var err error

		var emailChk string
		db.Db.QueryRow("SELECT email FROM profile WHERE email=$1", email).Scan(&emailChk)
		if len(emailChk) <= 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "user not found",
			})
			return
		}

		r, err := db.Db.Query("SELECT measure_date, ecg, heart_rate FROM health_data WHERE email=$1 ORDER BY measure_date DESC LIMIT $2", email, rowLimit)
		defer r.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": err.Error(),
			})
			return
		}

		var res []HealthData

		fmt.Println(res)

		for r.Next() {
			var hd HealthData
			r.Scan(&hd.Date, &hd.Ecg, &hd.Hr)
			res = append(res, hd)
		}

		c.JSON(http.StatusOK, gin.H{
			"email": email,
			"row":   len(res),
			"data":  res,
		})
	}
}
