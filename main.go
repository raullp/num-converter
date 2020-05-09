package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/raullp/num-converter/docs"
	"github.com/raullp/num-converter/error"
	"github.com/raullp/num-converter/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(error.ErrorReporter())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/api/v1/converter/number-to-words/:number", handler.ConvertToWords)
	return r
}

// @title Number Converter API
// @version 1.0
// @description This is a Number to English words converter.

// @contact.name Raul Lopez
// @contact.email raul.lp@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}
func main() {
	r := setupRouter()
	r.Run()
}
