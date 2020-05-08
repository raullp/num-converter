package error

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func ErrorReporter() gin.HandlerFunc {
	return jsonErrorReporterT(gin.ErrorTypeAny)
}

type errorResponse struct {
	Code   int    `json:"-"`
	Status string `json:"status"`
	Detail string `json:"detail"`
}

func jsonErrorReporterT(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(errType)

		if len(detectedErrors) == 0 {
			return
		}

		err := detectedErrors[0].Err

		response := &errorResponse{
			Code:   KindUnexpected,
			Detail: "Internal Server Error",
			Status: "error",
		}

		var e *Error
		if errors.As(err, &e) {
			response.Code = int(GetKind(e))
			response.Detail = e.Err.Error()
		}

		c.AbortWithStatusJSON(response.Code, response)
	}
}
