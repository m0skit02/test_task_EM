package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wb-task-L0/pkg/models"
)

// @Summary Create Subscription
// @Tags Subscription
// @Description Create a new subscription
// @ID create-subscription
// @Accept json
// @Produce json
// @Param input body models.Subscription true "Subscription Input"
// @Success 200 {object} map[string]string "returns subscription ID"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/subscription/ [post]
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

// @Summary Get All Subscriptions
// @Tags Subscription
// @Description Get all subscriptions
// @ID get-all-subscriptions
// @Accept json
// @Produce json
// @Success 200 {object} []models.Subscription
// @Failure 500 {object} errorResponse
// @Router /api/subscription/ [get]
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

// @Summary Get Subscription by ID
// @Tags Subscription
// @Description Get subscription by its ID
// @ID get-subscription-by-id
// @Accept json
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} models.Subscription
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/subscription/{id} [get]
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

// @Summary Delete Subscription
// @Tags Subscription
// @Description Delete a subscription by its ID
// @ID delete-subscription
// @Accept json
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/subscription/{id} [delete]
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

// @Summary Get Subscriptions Summary
// @Tags Subscription
// @Description Get total cost summary of subscriptions with optional filters
// @ID get-subscriptions-summary
// @Accept json
// @Produce json
// @Param user_id query string false "User ID"
// @Param service_name query string false "Service Name"
// @Param start query string true "Start date in format DD.MM.YYYY"
// @Param end query string true "End date in format DD.MM.YYYY"
// @Success 200 {object} map[string]interface{} "total_cost and currency"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/subscription/summary [get]
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
