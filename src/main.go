package main

import (
	"./error"
	"./handler"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(error.ErrorReporter())
	r.GET("/api/v1/converter/number-to-words/:number", handler.ConvertToWords)
	return r
}

func main() {
	r := setupRouter()
	r.Run()
}
