package main

import (
	"github.com/geraldofigueiredo/practice-go/testing/testable_code/controllers"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/ping", controllers.Ping)

	router.Run(":8080")
}
