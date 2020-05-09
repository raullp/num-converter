package handler

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/raullp/num-converter/error"
	"github.com/raullp/num-converter/model"
	"github.com/raullp/num-converter/numconv"
	log "github.com/sirupsen/logrus"
)

// ShowAccount godoc
// @Summary Convert number to words
// @Description Converts the number to English words
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param number path int true "Number"
// @Success 200 {object} model.NumToWords
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} error.errorResponse
// @Failure 500 {object} error.errorResponse
// @Router /converter/number-to-words/{number} [get]
func ConvertToWords(c *gin.Context) {
	const op error.Op = "convert.ConvertToWords"
	const badRequest error.Kind = error.KindBadRequest
	number := c.Param("number")
	
	re := regexp.MustCompile(`(0|([1-9]\d*))`)

	matches := re.MatchString(number)

	if !matches {
		c.Error(error.E(errors.New("number: is not a valid number"), op, badRequest, log.ErrorLevel))
		return
	}
	words := numconv.ConvertToWordsA(number)
	response := &model.NumToWords{
		Status:       "ok",
		NumInEnglish: words,
	}
	c.JSON(http.StatusOK, response)
}
