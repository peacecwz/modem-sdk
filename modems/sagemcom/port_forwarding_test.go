package sagemcom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPortForwarding(t *testing.T) {
	client := NewSagemcomClient(BaseAddress)
	isLoggedIn, err := client.Login(Password)
	assert.NoError(t, err)
	assert.True(t, isLoggedIn)

	resp, err := client.PortForwarding.GetPortForwardings()
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.PortForwarding.Rules, 2)

	defer client.Logout()
}
