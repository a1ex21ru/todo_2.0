package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	todo "github.com/alex21ru/todo_2.0"
)

// @Summary Создать новый список
// @Security ApiKeyAuth
// @Tags lists
// @Description Создать новый список задач
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body todo.TodoList true "Список задач"
// @Success 200 {object} int "id"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := getuserId(c)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "error for create list")
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []todo.TodoList
}

// @Summary Получить все списки
// @Security ApiKeyAuth
// @Tags lists
// @Description Получить все списки пользователя
// @ID get-all-lists
// @Produce  json
// @Success 200 {object} getAllListsResponse
// @Failure 500 {object} errorResponse
// @Router /api/lists [get]
func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getuserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})

}

// @Summary Получить список по id
// @Security ApiKeyAuth
// @Tags lists
// @Description Получить список задач по id
// @ID get-list-by-id
// @Produce  json
// @Param id path int true "ID списка"
// @Success 200 {object} todo.TodoList
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/lists/{id} [get]
func (h *Handler) getListById(c *gin.Context) {
	userId, err := getuserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logrus.Debugf("update query: %s", &err)
		return
	}

	c.JSON(http.StatusOK, list)
}

// @Summary Обновить список
// @Security ApiKeyAuth
// @Tags lists
// @Description Обновить данные списка задач
// @ID update-list
// @Accept  json
// @Produce  json
// @Param id path int true "ID списка"
// @Param input body todo.UpdateListInput true "Данные для обновления списка"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/lists/{id} [put]
func (h *Handler) updateList(c *gin.Context) {
	userId, err := getuserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input todo.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoList.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}

// @Summary Удалить список
// @Security ApiKeyAuth
// @Tags lists
// @Description Удалить список задач по id
// @ID delete-list
// @Produce  json
// @Param id path int true "ID списка"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/lists/{id} [delete]
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getuserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err_db := h.services.TodoList.Delete(userId, id)
	if err_db != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})
}
