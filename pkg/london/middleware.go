package london

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
				httpStatus := httpStatus(err)
				log.Errorf("request failed with error %s", err.Error())
				debug.PrintStack()
				c.XML(httpStatus, errorResponse{StatusCode: httpStatus, Error: err.Error()})
				c.Error(err)
				c.Abort()
			}
		}()

		c.Next()
	}
}

func httpStatus(err error) (httpStatus int) {
	switch err.(type) {

	case validationError:
		httpStatus = http.StatusBadRequest

		return

	case unAvailableBookingError:
		httpStatus = http.StatusUnauthorized

		return

	default:
		httpStatus = http.StatusInternalServerError

		return
	}
}
