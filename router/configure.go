package router

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", Home)
	r.GET("/flight/:id", GetFlightDataFromID)
	r.POST("/flight-status", GetFlightData)
	r.POST("/flight/:id/:status", UpdateFlightStatus)
	r.POST("/tokens", InsertToken)
	r.POST("/send-notifications", SendNotificationHandler)

	r.Run("localhost:8080")
}
