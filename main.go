package main

import (
	"net/http"

	db "github.com/SpandanMaj-git/Even-Booking/DB"
	"github.com/SpandanMaj-git/Even-Booking/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()
	r.GET("/events", GetEvents)
	r.POST("/events", createEvent)

	r.Run()
}

func GetEvents(c *gin.Context) {
	events := models.GetEvents()
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.BindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()
	c.JSON(http.StatusCreated, gin.H{"message": "event created successfully", "event": event})

}
