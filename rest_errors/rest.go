package rest_errors

import (
	"errors"
	"net/http"
)

type restErr struct {
	message string        `json:"message"`
	status  int           `json:"status"`
	error   string        `json:"error"`
	causes  []interface{} `json:"causes"`
}

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
	return restErr{
		message: message,
		status:  status,
		error:   err,
		causes:  causes,
	}
}

func NewBadRequestError(message string) RestErr {
	return restErr{
		message: message,
		status:  http.StatusBadRequest,
		error:   "bad_request",
	}
}

func NewNotFoundError(message string) RestErr {
	return restErr{
		message: message,
		status:  http.StatusNotFound,
		error:   "not found",
	}
}

func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		message: message,
		status:  http.StatusInternalServerError,
		error:   "Internal_server_error",
	}
	if err != nil {
		result.causes = append(result.causes, err.Error())
	}
	return result
}

func NewUnauthorizedError() RestErr {
	return restErr{
		message: "unable to retrieve information with the given access token",
		status:  http.StatusUnauthorized,
		error:   "unauthorized",
	}
}

func (e restErr) Message() string {
	return e.message
}

func (e restErr) Status() int {
	return e.status
}

func (e restErr) Error() string {
	return e.error
}

func (e restErr) Causes() []interface{} {
	return e.causes
}
