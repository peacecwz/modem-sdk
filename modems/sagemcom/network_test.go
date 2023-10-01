package sagemcom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNetworkHosts(t *testing.T) {
	client := NewSagemcomClient(BaseAddress)
	isLoggedIn, err := client.Login(Password)
	assert.NoError(t, err, "failed to login")
	assert.True(t, isLoggedIn, "failed to login")

	resp, err := client.Network.GetHosts(true)
	assert.NoError(t, err, "failed to get network hosts")
	assert.NotNil(t, resp, "failed to get network hosts")
	assert.NotEmpty(t, resp.Hosts.Hosts, "failed to get network hosts")
}
