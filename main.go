package main

import (
	"net/http"
	"uint_converter/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("static/*")

	r.GET("/",func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",gin.H{
			"title":"Welcome to my Uint Converter!",
		})
	})

	r.GET("/length", handler.ShwoLengthPage)


	r.GET("/weight",handler.ShowWeightPage)

	// start server
	r.Run(":8080")
}
