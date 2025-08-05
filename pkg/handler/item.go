package handler

import (
	"net/http"
	"strconv"

	todo "github.com/alex21ru/todo_2.0"
	"github.com/gin-gonic/gin"
)

// @Summary Создать новый элемент
// @Security ApiKeyAuth
// @Tags items
// @Description Создать новый элемент в списке
// @ID create-item
// @Accept  json
// @Produce  json
// @Param id path int true "ID списка"
// @Param input body todo.TodoItem true "Элемент"
// @Success 200 {object} int "id"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/lists/{id}/items [post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := getuserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id ")
	}

	var input todo.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "error for create item")
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id": id,
	})
}

// @Summary Получить все элементы списка
// @Security ApiKeyAuth
// @Tags items
// @Description Получить все элементы по id списка
// @ID get-all-items
// @Produce  json
// @Param id path int true "ID списка"
// @Success 200 {array} todo.TodoItem
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/lists/{id}/items [get]
func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getuserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id ")
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid list id ")
	}

	c.JSON(http.StatusOK, items)
}

// @Summary Получить элемент по id
// @Security ApiKeyAuth
// @Tags items
// @Description Получить элемент по id
// @ID get-item-by-id
// @Produce  json
// @Param id path int true "ID элемента"
// @Success 200 {object} todo.TodoItem
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/items/{id} [get]
func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getuserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id ")
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid list id ")
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Обновить элемент (заглушка)
// @Security ApiKeyAuth
// @Tags items
// @Description Обновить элемент (метод не реализован)
// @ID update-item-stub
// @Router /api/items/{id} [put]
func (h *Handler) updateItem(c *gin.Context) {

}

// @Summary Удалить элемент
// @Security ApiKeyAuth
// @Tags items
// @Description Удалить элемент по id
// @ID delete-item
// @Produce  json
// @Param id path int true "ID элемента"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/items/{id} [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getuserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id ")
	}

	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid list id ")
	}

	c.JSON(http.StatusOK, statusResponse{
		"ok",
	})
}

// @Summary Обновить элемент
// @Security ApiKeyAuth
// @Tags items
// @Description Обновить данные элемента
// @ID update-item
// @Accept  json
// @Produce  json
// @Param id path int true "ID элемента"
// @Param input body todo.UpdateItemInput true "Данные для обновления элемента"
// @Success 200 {object} statusResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/items/{id} [put]
func (h *Handler) UpdateItem(c *gin.Context) {
	userId, err := getuserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input todo.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoItem.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}
