package errs

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AppError struct {
	Code    int
	Status  string
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) error {
	return AppError{
		Code:    fiber.ErrNotFound.Code,
		Status:  fiber.ErrNotFound.Message,
		Message: message,
	}
}

func NewUnexpectedError() error {

	return AppError{
		Code:    fiber.ErrInternalServerError.Code,
		Status:  fiber.ErrInternalServerError.Message,
		Message: "something went wrong",
	}
}

func NewBadRequestError(message string) error {
	return AppError{
		Code:    fiber.ErrBadRequest.Code,
		Status:  fiber.ErrBadRequest.Message,
		Message: message,
	}
}

func IsErrForeignKeyViolated(err error) bool {
	return strings.Contains(err.Error(), "violates foreign key")
}
