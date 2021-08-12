package handler

/**
 * Authorization & Authentication handler
 */

import (
	"WebApplication"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Authorization
func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	// HTTP 400 - Incorrect data in http.request
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)

	// HTTP 503 - Service is unavailable
	if err != nil {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	// HTTP 201 - Created successfully
	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//Authentication
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	// HTTP 404 - User not found
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	authToken, err := h.services.Authorization.GenerateToken(input.Username, input.Password)

	// HTTP 503 - Service is unavailable
	if err != nil {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	// HTTP 200 - Successfully
	c.JSON(http.StatusOK, map[string]interface{}{
		"authorization token": authToken,
	})
}
