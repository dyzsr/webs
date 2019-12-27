package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.GET("/", view)
	app.POST("/upload", upload)
	app.POST("/download", download)
	app.DELETE("/delete", delete)
	app.Run()
}
