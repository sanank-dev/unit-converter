package main

import (
	"net/http"
	"uint_converter/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.Static("/static","./sttaic")

	r.Static("/static","./static")


	r.GET("/",func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",gin.H{
			"title":"Welcome to my Uint Converter!",
		})
	})

	r.GET("/length", handler.ShwoLengthPage)


	r.GET("/weight",handler.ShowWeightPage)


	r.GET("/temperature",handler.ShowTempraturePage)


	// start server
	r.Run(":8081")
}
