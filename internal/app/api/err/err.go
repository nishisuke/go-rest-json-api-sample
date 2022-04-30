package err

import (
	"fmt"
	"net/http"
)

type APIErr int

const (
	ErrInvalidContentType APIErr = iota + 1
	ErrQueryParamsExist
	ErrZeroPage
)

func (e APIErr) Error() string {
	return fmt.Sprintf("Error: %d", e)
}

func (e APIErr) HttpStatus() int {
	switch e {
	case ErrInvalidContentType, ErrQueryParamsExist, ErrZeroPage:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
