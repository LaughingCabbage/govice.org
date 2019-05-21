package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	templatesHome := os.Getenv("TEMPLATES_HOME")
	log.Println("Templates Home: " + templatesHome)

	if templatesHome != "" {
		r.LoadHTMLGlob(templatesHome + "/*")
	} else {
		r.LoadHTMLGlob("templates/*")
	}

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Govice | Home",
		})
	})
	r.Run()
}
