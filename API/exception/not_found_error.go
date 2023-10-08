package exception

// menyesuaikan error berupa interface yang memiliki atribute error
type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}
