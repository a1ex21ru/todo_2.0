package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type error_h struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statuscode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statuscode, error_h{Message: message})
}
