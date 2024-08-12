package routes

import (
	"github.com/EnesHarman/eventify/internal/controller"
	"github.com/EnesHarman/eventify/internal/repository"
	"github.com/EnesHarman/eventify/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	eventRepository := repository.NewEventRepository()
	eventService := service.NewEventService(eventRepository)
	eventController := controller.NewEventController(eventService)

	e.GET("/events", eventController.GetEvents)
	e.POST("/event/add", eventController.AddEvent)
}
