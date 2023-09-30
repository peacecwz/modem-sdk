package sagemcom

import (
	"encoding/json"
	"fmt"
	"github.com/peacecwz/modem-sdk/pkg/client"
)

type PortForwardingRule struct {
	ID   int `json:"id"`
	Rule struct {
		Enabled           bool   `json:"enabled"`
		ExternalStartPort int    `json:"externalStartPort"`
		ExternalEndPort   int    `json:"externalEndPort"`
		Protocol          string `json:"protocol"`
		LocalStartPort    int    `json:"localStartPort"`
		LocalEndPort      int    `json:"localEndPort"`
		LocalAddress      string `json:"localAddress"`
		ReadOnly          bool   `json:"readOnly"`
	}
}

type PortForwardingResponse struct {
	PortForwarding struct {
		Rules []PortForwardingRule `json:"rules"`
	} `json:"portforwarding"`
}

type PortForwardingClient struct {
	httpClient *client.Client
}

func NewSagemcomPortForwardingClient(client *client.Client) *PortForwardingClient {
	return &PortForwardingClient{
		httpClient: client,
	}
}

func (c *PortForwardingClient) GetPortForwardings() (*PortForwardingResponse, error) {
	resp, err := c.httpClient.GET("network/portforwarding", nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response PortForwardingResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}
