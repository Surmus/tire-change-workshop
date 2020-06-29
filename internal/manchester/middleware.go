package manchester

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
)

func errorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			r := recover()

			if err, ok := r.(error); ok {
				httpStatus, errorCode := httpStatus(err)
				log.Errorf("request failed with error %s", err.Error())
				debug.PrintStack()

				c.Error(err)
				c.AbortWithStatusJSON(httpStatus, errorResponse{Code: errorCode, Message: err.Error()})
			}
		}()

		c.Next()
	}
}

func httpStatus(err error) (httpStatus int, errorCode string) {
	if appErr, ok := err.(*tireChangeApplicationError); ok {
		switch appErr.code {
		case validationErrorCode:
			return http.StatusBadRequest, appErr.code

		case unAvailableTimeErrorCode:
			return http.StatusUnauthorized, appErr.code
		}
	}

	return http.StatusInternalServerError, "500"
}
