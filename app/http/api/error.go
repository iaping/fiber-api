package api

import "github.com/gofiber/fiber/v2"

var (
	ErrorServer    = NewWithStatus(fiber.StatusInternalServerError, 1, "sorry, an error occurred on the server")
	ErrorParameter = New(2, "invalid parameter")
)

type Error struct {
	Status int
	Code   int
	Text   string
}

func (e *Error) Error() string {
	return e.Text
}

func New(code int, text string) *Error {
	return NewWithStatus(fiber.StatusBadRequest, code, text)
}

func NewWithStatus(status, code int, text string) *Error {
	return &Error{status, code, text}
}
