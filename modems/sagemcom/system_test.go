package sagemcom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSystemInfo(t *testing.T) {
	client := NewSagemcomClient(BaseAddress)
	isLoggedIn, err := client.Login(Password)
	assert.NoError(t, err, "failed to login")
	assert.True(t, isLoggedIn, "failed to login")

	resp, err := client.System.GetSystemInfo()
	assert.NoError(t, err, "failed to get system info")
	assert.NotNilf(t, resp, "response is nil")
	assert.NotEmptyf(t, resp.Info.ModelName, "model name is empty")
	assert.NotEmptyf(t, resp.Info.SoftwareVersion, "software version is empty")
	assert.NotEmptyf(t, resp.Info.HardwareVersion, "hardware version is empty")

	defer client.Logout()
}

func TestGetSystemGatewayInfoProvisioning(t *testing.T) {
	client := NewSagemcomClient(BaseAddress)
	isLoggedIn, err := client.Login(Password)
	assert.NoError(t, err, "failed to login")
	assert.True(t, isLoggedIn, "failed to login")

	resp, err := client.System.GetSystemGatewayInfoProvisioning()
	assert.NoError(t, err, "failed to get system gateway info provisioning")
	assert.NotNilf(t, resp, "response is nil")
	assert.NotEmptyf(t, resp.Provisioning.Mode, "mode is empty")
	assert.NotEmptyf(t, resp.Provisioning.MacAddress, "mac address is empty")
	assert.NotEmptyf(t, resp.Provisioning.IPv4.Address, "ipv4 address is empty")
	assert.NotEmptyf(t, resp.Provisioning.IPv4.DefaultGateway, "ipv4 default gateway is empty")
	assert.NotEmptyf(t, resp.Provisioning.IPv4.DNSServers, "ipv4 dns servers is empty")
	assert.NotEmptyf(t, resp.Provisioning.IPv6.GlobalAddress, "ipv6 global address is empty")
	assert.NotEmptyf(t, resp.Provisioning.IPv6.LinkAddress, "ipv6 link address is empty")
	assert.NotEmptyf(t, resp.Provisioning.IPv6.DefaultGateway, "ipv6 default gateway is empty")
	assert.NotEmptyf(t, resp.Provisioning.IPv6.DNSServers, "ipv6 dns servers is empty")

	defer client.Logout()
}
