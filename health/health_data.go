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
		Type  string `json:"type"`
		Value int    `json:"value"`
	} `json:"data"`
}

type HealthData struct {
	Date     time.Time `json:"date" sql:"h.measure_date"`
	TypeName string    `json:"type" sql:"t.type_name"`
	Value    int       `json:"value" sql:"h.measure_value"`
}

type GetDataRequest struct {
	Email string `'json:"email"`
	Row   string `json:"row"`
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
		measureTypeId := 1
		if sdr.Data.Type == "ecg" {
			measureTypeId = 0
		}
		r, err := db.Db.Query("INSERT INTO health_data VALUES ($1, $2, $3, $4)", sdr.Email, measureTypeId, dataTimestamp, sdr.Data.Value)
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
		var gdr GetDataRequest
		if (c.BindJSON(&gdr)) != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error in server",
			})
		}
		// check if email is registered
		var err error

		var emailChk string
		db.Db.QueryRow("SELECT email FROM profile WHERE email=$1", gdr.Email).Scan(&emailChk)
		if len(emailChk) <= 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "user not found",
			})
			return
		}

		r, err := db.Db.Query(
			"SELECT h.measure_date, h.measure_value, t.type_name FROM health_data AS h, measure_type AS t WHERE h.email=$1 AND h.type_id=t.type_id ORDER BY measure_date DESC LIMIT $2",
			gdr.Email,
			gdr.Row,
		)
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
			r.Scan(&hd.Date, &hd.Value, &hd.TypeName)
			res = append(res, hd)
		}

		c.JSON(http.StatusOK, gin.H{
			"email": gdr.Email,
			"row":   len(res),
			"data":  res,
		})
	}
}
