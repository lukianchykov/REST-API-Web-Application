package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/**
 * Error response
 */

type error struct {
	Message string `json:"message"`
}

//Standard func for Error handling
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, error{message})
}
