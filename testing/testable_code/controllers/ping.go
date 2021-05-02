package controllers

import (
	"net/http"

	"github.com/geraldofigueiredo/practice-go/testing/testable_code/services"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	result, err := services.PingService.HandlePing()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.String(http.StatusOK, result+"\n")
}
