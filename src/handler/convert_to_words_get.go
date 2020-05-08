package handler

import (
	"errors"
	"net/http"
	"strconv"

	"../error"
	"../numconv"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ConvertToWords(c *gin.Context) {
	const op error.Op = "convert.ConvertToWords"
	const badRequest error.Kind = error.KindBadRequest
	number, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		c.Error(error.E(errors.New("number: is not a valid number"), op, badRequest, log.ErrorLevel))
		return
	}
	words := numconv.ConvertToWords(number)
	c.JSON(http.StatusOK, gin.H{
		"status":         "ok",
		"num_in_english": words,
	})
}
