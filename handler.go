package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	body = `
	{{define "T1"}}
		<html>
		<head>
			<title>Title</title>
		</head>
		<body>
			<p>Body</p>
			{{.}}
		</body>
		</html>
	{{end}}
	`
)

func view(c *gin.Context) {
	tmpl, err := template.New("test").Parse(body)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "T1", "pipeline")
	if err != nil {
		log.Fatal(err)
		return
	}

	c.String(http.StatusOK, "ok")
	// c.HTML(http.StatusOK, "test", nil)
}
