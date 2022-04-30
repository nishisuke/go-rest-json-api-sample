package err

import (
	"fmt"
	"net/http"
)

type APIError int

const (
	ErrInvalidContentType APIError = iota + 1
	ErrQueryParamsExist
	ErrZeroPage
)

func (e APIError) Error() string {
	return fmt.Sprintf("Error: %d", e)
}

func (e APIError) HTTPStatus() int {
	switch e {
	case ErrInvalidContentType, ErrQueryParamsExist, ErrZeroPage:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
