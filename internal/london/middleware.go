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
				c.XML(httpStatus, errorResponse{StatusCode: httpStatus, Error: err.Error()})
				_ = c.Error(err)
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

		log.Debugf("request failed with error: %v", err)

		return

	case invalidTireChangeTimesPeriodError:
		httpStatus = http.StatusBadRequest

		log.Debugf("request failed with error: %v", err)

		return

	case unAvailableBookingError:
		httpStatus = http.StatusUnauthorized

		log.Debugf("request failed with error: %v", err)

		return

	default:
		httpStatus = http.StatusInternalServerError

		log.Errorf("request failed with error: %v", err)
		debug.PrintStack()

		return
	}
}
