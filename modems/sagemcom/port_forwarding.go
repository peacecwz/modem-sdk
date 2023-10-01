package sagemcom

import (
	"encoding/json"
	"fmt"
	"github.com/peacecwz/modem-sdk/pkg/client"
)

type PortForwardingRule struct {
	ID   int `json:"id"`
	Rule struct {
		Enabled           bool   `json:"enable"`
		ExternalStartPort int    `json:"externalStartPort"`
		ExternalEndPort   int    `json:"externalEndPort"`
		Protocol          string `json:"protocol"`
		LocalStartPort    int    `json:"localStartPort"`
		LocalEndPort      int    `json:"localEndPort"`
		LocalAddress      string `json:"localAddress"`
		ReadOnly          bool   `json:"readOnly"`
	} `json:"rule"`
}

type CreatePortForwardingResponse struct {
	Created struct {
		ID  int    `json:"id"`
		URI string `json:"uri"`
	} `json:"created"`
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

func (c *PortForwardingClient) CreatePortForwarding(req PortForwardingRule) (*CreatePortForwardingResponse, error) {
	resp, err := c.httpClient.POST("network/portforwarding", nil, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response CreatePortForwardingResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}

func (c *PortForwardingClient) UpdatePortForwarding(req PortForwardingRule) (bool, error) {
	resp, err := c.httpClient.PUT(fmt.Sprintf("network/portforwarding/%d", req.ID), nil, req)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	return resp.StatusCode == 204, nil
}

func (c *PortForwardingClient) DeletePortForwarding(id int) (bool, error) {
	resp, err := c.httpClient.DELETE(fmt.Sprintf("network/portforwarding/%d", id), nil)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	return resp.StatusCode == 204, nil
}
