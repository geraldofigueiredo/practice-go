package controllers

import (
	"net/http"

	"github.com/geraldofigueiredo/practice-go/testing/testable_code/services"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	result := services.HandlePing()
	c.String(http.StatusOK, result)
}
