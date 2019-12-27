package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	filenames = map[string]string{
		"index": "templates/index.html",
	}

	body string
)

func init() {
	data, err := ioutil.ReadFile("templates/index.html")
	if err != nil {
		panic(err)
	}
	body = string(data)

	log.Printf("template <%s> loaded\n", filenames["index"])
}

func view(c *gin.Context) {
	tmpl, err := template.New("index").Parse(body)
	if err != nil {
		log.Fatal(err)
		c.String(http.StatusInternalServerError, "error")
		return
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "T1", "pipeline")
	if err != nil {
		log.Fatal(err)
		c.String(http.StatusInternalServerError, "error")
		return
	}

	c.HTML(http.StatusOK, "index", "pipeline")
}

func upload(c *gin.Context) {

}

func download(c *gin.Context) {

}

func delete(c *gin.Context) {

}
