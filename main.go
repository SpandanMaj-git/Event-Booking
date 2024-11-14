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

func createEvent(c *gin.Context) {
	var event models.Event
	c.BindJSON(&event)
}
