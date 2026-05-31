package handler

import (
	"net/http"
	"uint_converter/converter"

	"github.com/gin-gonic/gin"
)


func ShowWeightPage(c *gin.Context)  {
	c.HTML(http.StatusOK,"weight.html",nil)
}


func ConvertWeight(c *gin.Context)  {
	valueStr := c.PostForm("value")
	from := c.PostForm("from")
	to := c.PostForm("to")

	val,err := converter.WeightConverter(valueStr,from,to)

	if err != nil {
        c.HTML(http.StatusBadRequest, "weight.html", gin.H{
            "result": "❌ Invalid input. Please enter a valid number.",
        })
        return
    }
	c.HTML(http.StatusOK, "weight.html", gin.H{
        "result": val,
    })
}