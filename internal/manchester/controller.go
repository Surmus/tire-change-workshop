package manchester

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const v1Path = "/api/v1"

type controller struct {
	service *tireChangeTimesService
}

func registerController(router *gin.Engine, service *tireChangeTimesService) {
	c := &controller{service: service}

	router.GET(v1Path+"/tire-change-times", c.getTireChangeTimes)
	router.PUT(v1Path+"/tire-change-times/:id/book", c.putTireChangeBooking)
}

// getTireChangeTimes godoc
// @Summary List of available tire change times
// @Accept json
// @Produce json
// @Param amount query integer false "amount of tire change times per page"
// @Param page query integer false "The number of pages to skip before starting to collect the result set"
// @Param from query string false "search tire change times from date" Format(date) default(2006-01-02)
// @Success 200 {object} tireChangeTimesResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /tire-change-times [get]
func (c *controller) getTireChangeTimes(ctx *gin.Context) {
	var query tireChangeTimesSearchQuery

	if err := ctx.ShouldBind(&query); err != nil {
		panic(newValidationError(err))
	}

	ctx.JSON(http.StatusOK, c.service.get(&query))
}

// putTireChangeBooking godoc
// @Summary Book tire change time
// @Accept json
// @Produce json
// @Param id path integer true "available tire change time ID"
// @Param body body tireChangeBookingRequest true "Request body"
// @Success 200 {object} tireChangeTimeResponse
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /tire-change-times/{id}/book [put]
func (c *controller) putTireChangeBooking(ctx *gin.Context) {
	var uri tireChangeBookingURI
	var request tireChangeBookingRequest

	if err := ctx.ShouldBindUri(&uri); err != nil {
		panic(newValidationError(err))
	} else if err := ctx.ShouldBindJSON(&request); err != nil {
		panic(newValidationError(err))
	}

	response, err := c.service.book(uri.ID, request.ContactInformation)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, response)
}
