package mr

// ErrInvalidURL is returned when the given connection URL could not be parsed
type ErrInvalidURL string

func (e ErrInvalidURL) Error() string {
	return "Invalid MongoDB URL: " + string(e)
}
