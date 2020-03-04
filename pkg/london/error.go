package london

type validationError struct {
	error
}

type unAvailableBookingError struct {
	error string
}

func newUnAvailableBookingError() unAvailableBookingError {
	return unAvailableBookingError{error: "tire change time is unavailable"}
}

func (e unAvailableBookingError) Error() string {
	return e.error
}
