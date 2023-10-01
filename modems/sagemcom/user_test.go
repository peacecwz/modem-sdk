package sagemcom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogin(t *testing.T) {
	client := NewSagemcomClient(BaseAddress)
	isLoggedIn, err := client.Login(Password)
	assert.NoError(t, err, "failed to login")
	assert.True(t, isLoggedIn, "failed to login")
	assert.NotEmptyf(t, client.LoggedSession.Created.Token, "failed to login")
}

func TestLogout(t *testing.T) {
	client := NewSagemcomClient(BaseAddress)
	isLoggedIn, err := client.Login(Password)
	assert.NoError(t, err, "failed to login")
	assert.True(t, isLoggedIn, "failed to login")

	isLoggedOut, err := client.Logout()
	assert.NoError(t, err, "failed to logout")
	assert.True(t, isLoggedOut, "failed to logout")
}
