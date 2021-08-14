package handler

/**
 * Items handler
 */

import (
	todo "WebApplication"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	// HTTP 400 - Incorrect data in http.request
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input todo.TodoItem
	// HTTP 400 - Incorrect data in http.request
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(userId, listId, input)
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

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	// HTTP 400 - Incorrect data in http.request
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)
	// HTTP 503 - Service is unavailable
	if err != nil {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	// HTTP 200 - Successfully
	c.JSON(http.StatusOK, getAllItemsResponse{Data: items})
}

func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	// HTTP 400 - Incorrect data in http.request
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)
	// HTTP 503 - Service is unavailable
	if err != nil {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	// HTTP 200 - Successfully
	c.JSON(http.StatusOK, item)
}

func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	// HTTP 400 - Incorrect data in http.request
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateItemInput
	// HTTP 400 - Incorrect data in http.request
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// HTTP 503 - Service is unavailable
	if err := h.services.TodoItem.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	// HTTP 200 - Successfully
	c.JSON(http.StatusOK, statusResponse{"Updated Successfully"})
}

func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	// HTTP 503 - Service is unavailable
	if err != nil {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	// HTTP 400 - Incorrect data in http.request
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	// HTTP 503 - Service is unavailable
	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	// HTTP 200 - Successfully
	c.JSON(http.StatusOK, statusResponse{"Deleted Successfully"})
}
