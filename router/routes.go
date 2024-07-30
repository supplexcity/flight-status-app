package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"test/db"
	"test/firebase"

	"github.com/gin-gonic/gin"
)

func GetFlightData(c *gin.Context) {
	id := c.PostForm("flight_id")
	fmt.Printf("Request for flight ID = %s\n", id)
	flightID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Error in converting string to int")
		c.IndentedJSON(http.StatusBadRequest, id)
		return
	}
	f := db.GetFlightDetails(flightID)
	c.IndentedJSON(http.StatusOK, f)
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func AddName(c *gin.Context) {
	f := c.PostForm("first_name")
	l := c.PostForm("last_name")
	fmt.Printf("Name = %s %s\n", f, l)
}

func UpdateFlightStatus(c *gin.Context) {
	id := c.Param("id")
	flightID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Error in converting string to int")
		c.IndentedJSON(http.StatusBadRequest, id)
		return
	}
	newStatus := c.Param("status")
	db.UpdateFlightStatus(flightID, newStatus)
}

func GetFlightDataFromID(c *gin.Context) {
	// getting data
	id := c.Param("id")
	flightID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Error in converting string to int")
		c.IndentedJSON(http.StatusBadRequest, id)
		return
	}
	f := db.GetFlightDetails(flightID)
	c.IndentedJSON(http.StatusOK, f)
}

func InsertToken(c *gin.Context) {
	var token db.NotificationToken
	err := json.NewDecoder(c.Request.Body).Decode(&token)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, token)
		return
	}

	// Insert the token into the database
	err = db.InsertToken(token, c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, token)
		return
	}
}

func SendNotificationHandler(c *gin.Context) {
	var message db.Message
	err := json.NewDecoder(c.Request.Body).Decode(&message)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "")
		return
	}
	tokens, err := db.GetTokens(message.UserId, c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	err = firebase.SendNotification(c, tokens, message.UserId, message.Text)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
}
