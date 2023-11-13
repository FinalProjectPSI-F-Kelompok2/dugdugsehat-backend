package main

import (
	"github.com/FinalProjectPSI-F-Kelompok2/dugdugsehat-backend/profile"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// User Login
	router.POST("/login", profile.Login())
	router.Run(":8080")
}
