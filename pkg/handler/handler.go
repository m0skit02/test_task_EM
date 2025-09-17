package handler

import (
	"github.com/gin-gonic/gin"
	"wb-task-L0/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler { return &Handler{services: services} }

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		subscription := api.Group("/subscription")
		{
			subscription.POST("/", h.createSubscription)
			subscription.GET("/", h.getAllSubscriptions)
			subscription.GET("/:ID", h.getByIDSubscriptions)
			subscription.DELETE("/:ID", h.deleteSubscription)
			subscription.GET("/summary", h.getSubscriptionsSummary)
		}
	}

	return router
}
