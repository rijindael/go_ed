package logic

import (
	"net_pro/handlers"

	"github.com/gin-gonic/gin"
)

func Run_back() {
	router := gin.Default()
	router.GET("/", handlers.Getting)
	router.POST("/", handlers.Posting)
	router.Run(":3000")
}
