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
}

func (h *Handler) singIn(c *gin.Context) {
}
