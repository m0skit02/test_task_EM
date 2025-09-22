package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "wb-task-EM/docs"

	"wb-task-EM/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler { return &Handler{services: services} }

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		subscription := api.Group("/subscription")
		{
			subscription.POST("/", h.createSubscription)
			subscription.GET("/", h.getAllSubscriptions)
			subscription.GET("/:id", h.getByIDSubscriptions)
			subscription.DELETE("/:id", h.deleteSubscription)
			subscription.GET("/summary", h.getSubscriptionsSummary)
		}
	}

	return router
}
