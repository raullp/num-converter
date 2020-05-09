package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raullp/num-converter/error"
	"github.com/raullp/num-converter/model"
	"github.com/raullp/num-converter/numconv"
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
	response := &model.NumToWords{
		Status:       "ok",
		NumInEnglish: words,
	}
	c.JSON(http.StatusOK, response)
}
