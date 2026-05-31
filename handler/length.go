package handler

import (
	"net/http"
	"uint_converter/converter"

	"github.com/gin-gonic/gin"
)



func ShwoLengthPage(c *gin.Context)  {
	c.HTML(http.StatusOK, "length.html", nil)
}

func ConvertLength(c *gin.Context) {
	valueStr := c.PostForm("value")
	from := c.PostForm("from")
	to := c.PostForm("to")

	val, err := converter.LenthConverter(valueStr, from, to)

	if err != nil {
		// If strconv.ParseFloat failed inside the converter, handle it here where 'c' is available
		c.String(http.StatusBadRequest, "invalid number format")
		return
	}

	c.HTML(http.StatusOK, "length.html", gin.H{

		"result": val,
	})
}
