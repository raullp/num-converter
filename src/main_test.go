package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Response struct {
	NumInEnglish string `json:"num_in_english"`
	Status       string `json:"status"`
}

func TestConvertNumberToWords(t *testing.T) {
	router := setupRouter()

	words := "twelve thousand three hundred forty-five"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/converter/number-to-words/12345", nil)
	router.ServeHTTP(w, req)

	response := Response{}
	json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "success", response.Status)
	assert.Equal(t, words, response.NumInEnglish)
}
