package api

import "github.com/gofiber/fiber/v2"

var (
	ErrorServer    = NewErrorWithStatus(fiber.StatusInternalServerError, 1, "sorry, an error occurred on the server")
	ErrorParameter = NewError(2, "invalid parameter")
)

type Error struct {
	Status int
	Code   int
	Text   string
}

func (e *Error) Error() string {
	return e.Text
}

func NewError(code int, text string) *Error {
	return NewErrorWithStatus(fiber.StatusBadRequest, code, text)
}

func NewErrorWithStatus(status, code int, text string) *Error {
	return &Error{status, code, text}
}
