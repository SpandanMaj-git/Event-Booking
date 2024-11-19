package main

import (
	"net/http"

	"github.com/SpandanMaj-git/Even-Booking/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/events", GetEvents)
	r.POST("/events", createEvent)

	r.Run()
}

func GetEvents(c *gin.Context) {
	events := models.GetEvents()
	c.JSON(http.StatusOK, events)
}

// func createEvent(c *gin.Context) {
// 	var event models.Event
// 	err := c.BindJSON(&event)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Invalid request",
// 		})
// 		return
// 	}

// 	event.ID = 1
// 	event.UserID = 1
// 	c.JSON(http.StatusCreated, gin.H{"message": "event created successfully", "event": event})

// }

func createEvent(c *gin.Context) {
	var event models.Event

	// Bind JSON and validate
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"details": err.Error(),
		})
		return
	}

	// Set ID and UserID manually (or generate dynamically)
	event.ID = len(models.GetEvents()) + 1
	event.UserID = 1

	// Save the event
	event.Save()

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})
}
