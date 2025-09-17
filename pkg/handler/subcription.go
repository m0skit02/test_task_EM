package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wb-task-L0/pkg/models"
)

func (h *Handler) createSubscription(c *gin.Context) {
	var input models.Subscription
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	sub, err := h.services.Subscription.Create(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": sub.ID,
	})
}

// GET /subscriptions
func (h *Handler) getAllSubscriptions(c *gin.Context) {
	subs, err := h.services.Subscription.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": subs,
	})
}

// GET /subscriptions/:id
func (h *Handler) getByIDSubscriptions(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	sub, err := h.services.Subscription.GetByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, sub)
}

// DELETE /subscriptions/:id
func (h *Handler) deleteSubscription(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err := h.services.Subscription.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) getSubscriptionsSummary(c *gin.Context) {
	userID := c.Query("user_id")
	serviceName := c.Query("service_name")
	start := c.Query("start")
	end := c.Query("end")

	if start == "" || end == "" {
		newErrorResponse(c, http.StatusBadRequest, "start and end query params are required (MM-YYYY)")
		return
	}

	total, err := h.services.Subscription.GetTotalCost(userID, serviceName, start, end)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_cost": total,
		"currency":   "RUB",
	})
}
