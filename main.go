package main

import (
	"main/handler"
	"main/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.MyLogger())

	// 웹 소켓 핸들러 등록
	r.GET("/ws", handler.HandleWebSocket)

	r.Run(":8081")
}
