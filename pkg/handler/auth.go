package handler

import (
	"net/http"

	todo "github.com/alex21ru/todo_2.0"
	"github.com/gin-gonic/gin"
)

// @Summary Регистрация пользователя
// @Tags auth
// @Description Создать нового пользователя
// @ID create-user
// @Accept  json
// @Produce  json
// @Param input body todo.User true "Пользователь"
// @Success 200 {object} string "id"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		// logrus.Errorf("error of binding: %s", err.Error())
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary Авторизация пользователя
// @Tags auth
// @Description Вход пользователя и получение JWT
// @ID login-user
// @Accept  json
// @Produce  json
// @Param input body signInInput true "Данные для входа"
// @Success 200 {object} string "token"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		// logrus.Errorf("error of binding: %s", err.Error())
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"token": token,
	})
}
