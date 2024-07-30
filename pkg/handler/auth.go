package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todo "github.com/maximka200/NotesWebApp"
	"github.com/sirupsen/logrus"
)

func (h *Handler) singUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		logrus.Error(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) singIn(c *gin.Context) {
}
