package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

/**
 * Middleware projection for api HTTP group
 */

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	// HTTP 401 - User isn`t authorized. Need authentication
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty authorization header")
		return
	}
	// HTTP 401 - User isn`t authorized. Need authentication
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid authorization header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	//parse token
	// HTTP 401 - User isn`t authorized. Need authentication
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)

}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	// HTTP 404 - User not found
	if !ok {
		newErrorResponse(c, http.StatusNotFound, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	// HTTP 404 - User not found
	if !ok {
		newErrorResponse(c, http.StatusNotFound, "user id is of invalid type")
		return 0, errors.New("user id not found")
	}

	return idInt, nil
}
