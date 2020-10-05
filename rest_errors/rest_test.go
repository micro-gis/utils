package rest_errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	err:= NewInternalServerError("any message", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "any message", err.Message)
	assert.EqualValues(t, "Internal_server_error", err.Err)
	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes) )
	assert.EqualValues(t, "database error", err.Causes[0])
}

func TestNewBadRequestError(t *testing.T) {
//TODO : implement
}

func TestNewNotFoundError(t *testing.T) {
	//TODO : implement
}

func TestNewError(t *testing.T) {
	//TODO : implement
}