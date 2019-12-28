package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", view)
	router.GET("/media/:filename", download)
	router.POST("/upload", upload)
	router.Run(":9999")
}
