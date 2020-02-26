package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNotFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "Expected user not be found")
	assert.NotNil(t, err, "Expected to receive a proper error")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
}

func TestGetUserFound(t *testing.T) {
	user, err := GetUser(1)

	assert.NotNil(t, user, "Expected user to be fond")
	assert.Nil(t, err, "Expected no error to be received")
}
