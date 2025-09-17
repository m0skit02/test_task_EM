package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// pkg/handler/subscription.go
func (h *Handler) getSubscriptionsSummary(c *gin.Context) {
	userID := c.Query("user_id")
	serviceName := c.Query("service_name")
	start := c.Query("start")
	end := c.Query("end")

	if start == "" || end == "" {
		newErrorResponse(c, http.StatusBadRequest, "start and end query params are required (YYYY-MM-DD)")
		return
	}

	total, err := h.services.Subscription.GetTotalCost(userID, serviceName, start, end)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"total_cost": total,
	})
}
