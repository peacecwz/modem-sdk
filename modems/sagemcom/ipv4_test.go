package sagemcom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetIPV4Info(t *testing.T) {
	client := NewSagemcomClient(BaseAddress)
	isLoggedIn, err := client.Login(Password)
	assert.NoError(t, err, "failed to login")
	assert.True(t, isLoggedIn, "failed to login")

	resp, err := client.IPV4.GetInfo()
	assert.NoError(t, err, "failed to get ipv4 info")
	assert.NotNil(t, resp, "response is nil")
	assert.NotEmpty(t, resp.Info.LanIPAddress, "lan ip address is empty")
	assert.NotEmpty(t, resp.Info.LanSubnet, "lan subnet is empty")
	assert.NotEmpty(t, resp.Info.LanSubnetMask, "lan subnet mask is empty")

	defer client.Logout()
}
