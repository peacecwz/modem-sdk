package sagemcom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	BaseAddress = "http://192.168.178.1/rest/v1/"
	Password    = "Admin1224"
)

func TestClientLogin(t *testing.T) {
	client := NewSagemcomClient(BaseAddress)
	isLoggedIn, err := client.Login(Password)

	defer client.Logout()

	assert.NoError(t, err, "should not return error")
	assert.True(t, isLoggedIn, "should be logged in")
}

func TestClientLogout(t *testing.T) {
	client := NewSagemcomClient(BaseAddress)
	_, err := client.Login(Password)
	assert.NoError(t, err, "should not return error")

	isLoggedOut, err := client.Logout()
	assert.NoError(t, err, "should not return error")
	assert.True(t, isLoggedOut, "should be logged out")
}
