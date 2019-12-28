package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func view(c *gin.Context) {
	pageData, err := NewPageData()
	if err != nil {
		log.Fatal(err)
		c.String(http.StatusInternalServerError, "internal error")
		return
	}
	c.HTML(http.StatusOK, "index.tmpl", pageData)
}

func download(c *gin.Context) {
	filename := c.Param("filename")
	c.File("media/" + filename)
}

func upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Fatal(err)
		c.String(http.StatusBadRequest, "get form error")
		return
	}

	filename := "media/" + filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		log.Fatal(err)
		c.String(http.StatusInternalServerError, "save file error")
		return
	}

	c.String(http.StatusOK, "success")
}
