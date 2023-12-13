package main

import (
	"log"

	"github.com/FinalProjectPSI-F-Kelompok2/dugdugsehat-backend/health"
	"github.com/FinalProjectPSI-F-Kelompok2/dugdugsehat-backend/model"
	"github.com/FinalProjectPSI-F-Kelompok2/dugdugsehat-backend/profile"
	"github.com/gin-gonic/gin"
)

func main() {
	// DB Init
	var db model.DbCon
	err := db.ConnectDb()
	if err != nil {
		log.Fatalln(err.Error())
	}
	router := gin.Default()

	// User Login
	router.POST("/login", profile.Login(&db))
	router.POST("/register", profile.Register(&db))

	// User Profile
	router.POST("/profile/get", profile.GetProfile(&db))
	router.POST("/profile/update", profile.UpdateProfile(&db))

	// Health Data
	router.POST("/health/get", health.GetHealthData(&db))
	router.POST("/health/save", health.SaveHealthData(&db))
	router.Run(":8080")

	db.CloseDb()
}
