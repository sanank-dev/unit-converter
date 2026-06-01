package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShwoLengthPage(c *gin.Context) {
	c.HTML(http.StatusOK, "length.html", nil)
}


func ShowWeightPage(c *gin.Context)  {
	c.HTML(http.StatusOK,"weight.html",nil)
}

func ShowTempraturePage(c *gin.Context)  {
	c.HTML(http.StatusOK,"temprature.html",nil)
}
