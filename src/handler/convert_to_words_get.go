package handler

import (
	"net/http"
	"strconv"

	"../numconv"
	"github.com/gin-gonic/gin"
)

func ConvertToWords(c *gin.Context) {
	number, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		c.Error(err) // TODO: wrap this error as an Invalid Parameter
		return
	}
	words := numconv.ConvertToWords(number)
	c.JSON(http.StatusOK, gin.H{
		"status":         "success",
		"num_in_english": words,
	})
}
