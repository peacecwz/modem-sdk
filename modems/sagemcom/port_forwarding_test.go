package sagemcom

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"math/rand"
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

func TestCreatePortForwarding(t *testing.T) {
	client := NewSagemcomClient(BaseAddress)
	isLoggedIn, err := client.Login(Password)
	assert.NoError(t, err)
	assert.True(t, isLoggedIn)

	port := rand.Intn(65535-1024) + 1024

	rule := PortForwardingRule{}
	rule.Rule.Enabled = true
	rule.Rule.ExternalStartPort = port
	rule.Rule.ExternalEndPort = port
	rule.Rule.Protocol = "tcp"
	rule.Rule.LocalStartPort = port
	rule.Rule.LocalEndPort = port
	rule.Rule.LocalAddress = "192.168.178.161"
	rule.Rule.ReadOnly = false

	resp, err := client.PortForwarding.CreatePortForwarding(rule)
	assert.NoError(t, err, "failed to create port forwarding rule")
	assert.NotNil(t, resp, "response is nil")
	assert.Greater(t, resp.Created.ID, 0, "id is not greater than 0")
	logrus.Infof("Created port forwarding rule: %v", resp)

	client.PortForwarding.DeletePortForwarding(resp.Created.ID)
	defer client.Logout()
}

func TestUpdatePortForwarding(t *testing.T) {
	client := NewSagemcomClient(BaseAddress)
	isLoggedIn, err := client.Login(Password)
	assert.NoError(t, err)
	assert.True(t, isLoggedIn)

	port := rand.Intn(65535-1024) + 1024

	rule := PortForwardingRule{}
	rule.Rule.Enabled = true
	rule.Rule.ExternalStartPort = port
	rule.Rule.ExternalEndPort = port
	rule.Rule.Protocol = "tcp"
	rule.Rule.LocalStartPort = port
	rule.Rule.LocalEndPort = port
	rule.Rule.LocalAddress = "192.168.178.161"
	rule.Rule.ReadOnly = false

	resp, err := client.PortForwarding.CreatePortForwarding(rule)
	assert.NoError(t, err, "failed to create port forwarding rule")
	assert.NotNil(t, resp, "response is nil")
	assert.Greater(t, resp.Created.ID, 0, "id is not greater than 0")
	logrus.Infof("Created port forwarding rule: %v", resp)

	rule.Rule.Enabled = false
	rule.ID = resp.Created.ID
	isUpdated, err := client.PortForwarding.UpdatePortForwarding(rule)
	assert.NoError(t, err, "failed to update port forwarding rule")
	assert.True(t, isUpdated, "failed to update port forwarding rule")

	rules, err := client.PortForwarding.GetPortForwardings()
	assert.NoError(t, err, "failed to get port forwarding rules")
	assert.NotNil(t, rules, "response is nil")
	for _, r := range rules.PortForwarding.Rules {
		if r.ID == rule.ID {
			assert.False(t, r.Rule.Enabled, "rule is not disabled")
		}
	}

	//client.PortForwarding.DeletePortForwarding(resp.Created.ID)
	defer client.Logout()
}

func TestDeletePortForwarding(t *testing.T) {
	client := NewSagemcomClient(BaseAddress)
	isLoggedIn, err := client.Login(Password)
	assert.NoError(t, err)
	assert.True(t, isLoggedIn)

	port := rand.Intn(65535-1024) + 1024

	rule := PortForwardingRule{}
	rule.Rule.Enabled = true
	rule.Rule.ExternalStartPort = port
	rule.Rule.ExternalEndPort = port
	rule.Rule.Protocol = "tcp"
	rule.Rule.LocalStartPort = port
	rule.Rule.LocalEndPort = port
	rule.Rule.LocalAddress = "192.168.178.161"
	rule.Rule.ReadOnly = false

	resp, err := client.PortForwarding.CreatePortForwarding(rule)
	assert.NoError(t, err, "failed to create port forwarding rule")
	assert.NotNil(t, resp, "response is nil")
	assert.Greater(t, resp.Created.ID, 0, "id is not greater than 0")
	logrus.Infof("Created port forwarding rule: %v", resp)

	isDeleted, err := client.PortForwarding.DeletePortForwarding(resp.Created.ID)
	assert.NoError(t, err, "failed to delete port forwarding rule")
	assert.True(t, isDeleted, "failed to delete port forwarding rule")

	defer client.Logout()
}
