package handler

import (
	todo "WebApplication"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
 * Lists handler
 */

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input todo.TodoList
	// HTTP 400 - Incorrect data in http.request
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	// HTTP 503 - Service is unavailable
	if err != nil {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	// HTTP 200 - Successfully
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	// HTTP 503 - Service is unavailable
	if err != nil {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	// HTTP 200 - Successfully
	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	// HTTP 400 - Incorrect data in http.request
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetById(userId, id)
	// HTTP 503 - Service is unavailable
	if err != nil {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	// HTTP 200 - Successfully
	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
