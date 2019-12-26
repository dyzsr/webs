package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleInfo(c *gin.Context) {
	log.Println("info")
	c.String(http.StatusOK, "info")
}
