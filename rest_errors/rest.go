package rest_errors

import (
	"fmt"
	"errors"
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Err     string `json:"error"`
	Causes []interface{} `json:"causes"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func (e *RestErr) Error() string {
	return fmt.Sprintf("REST ERROR : %s:%d", e.Message, e.Status)
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Err:     "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Err:     "not found",
	}
}

func NewInternalServerError(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Err:     "Internal_server_error",

	}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}
