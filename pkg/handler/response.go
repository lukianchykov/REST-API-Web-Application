package handler

import (
	todo "WebApplication"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/**
 * Class represents error response
 */

type errorResponse struct {
	Message string `json:"message"`
}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

type getAllItemsResponse struct {
	Data []todo.TodoItem `json:"data"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
