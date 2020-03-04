package manchester

const (
	validationErrorCode      = "11"
	unAvailableTimeErrorCode = "22"
)

type tireChangeApplicationError struct {
	code  string
	error string
}

func (e tireChangeApplicationError) Error() string {
	return e.error
}

func newValidationError(cause error) *tireChangeApplicationError {
	return &tireChangeApplicationError{code: validationErrorCode, error: cause.Error()}
}

func newUnAvailableBookingError() *tireChangeApplicationError {
	return &tireChangeApplicationError{code: unAvailableTimeErrorCode, error: "tire change time is unavailable"}
}
