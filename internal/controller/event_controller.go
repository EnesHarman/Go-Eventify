package controller

import (
	"github.com/EnesHarman/eventify/internal/model"
	"github.com/EnesHarman/eventify/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EventController struct {
	eventService service.EventService
}

func NewEventController(eventService service.EventService) *EventController {
	return &EventController{
		eventService: eventService,
	}
}

func (ctrl EventController) AddEvent(c *gin.Context) {
	var event model.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.eventService.InsertEvent(event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, event)
}

func (ctrl EventController) GetEvents(c *gin.Context) {
	page, size := parsePageAndSize(c)
	events, err := ctrl.eventService.GetEvents(page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)

}

func parsePageAndSize(c *gin.Context) (int64, int64) {
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("size", "10")

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1 // or handle the error as needed
	}

	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10 // or handle the error as needed
	}
	return page, size
}
