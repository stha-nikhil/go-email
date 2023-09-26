package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stha-nikhil/go-email/api"
)

func main() {
	router := gin.Default()
	router.POST("/send-email", api.SendEmail)
	router.Run()
}
