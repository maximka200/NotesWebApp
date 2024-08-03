package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todo "github.com/maximka200/NotesWebApp"
)

func (h *Handler) singUp(c *gin.Context) {
	var input todo.User
	// десериализация данных из контекста c в input
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// создаем пользователя используя input
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// отправка отчета клиенту
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// структура для входа
type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) singIn(c *gin.Context) {
	var input signInInput
	// десериализация данных из контекста c в input
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// создаем токен используя input
	token, err := h.services.Authorization.CreateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// отправка отчета клиенту
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
