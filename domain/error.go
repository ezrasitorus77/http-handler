package domain

import (
	"fmt"
)

type (
	Error struct {
		HttpStatus   int
		ResponseCode string
		Reason       string
	}
)

func (e *Error) Error() string {
	return fmt.Sprintf(e.Reason)
}
