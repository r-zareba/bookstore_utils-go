package rest_errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestBadRequestError(t *testing.T) {
	err := InternalServerError("Some Error message", errors.New("some error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Internal Server Error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "some error", err.Causes[0])
}

func TestInternalServerError(t *testing.T) {

}

func TestNotFoundError(t *testing.T) {

}