package sagemcom

import (
	"encoding/json"
	"fmt"
	"github.com/peacecwz/modem-sdk/pkg/client"
)

type NetworkHostResponse struct {
	Hosts struct {
		Hosts []struct {
			MacAddress string `json:"macAddress"`
			Config     struct {
				Connected  bool    `json:"connected"`
				DeviceName string  `json:"deviceName"`
				DeviceType string  `json:"deviceType"`
				Hostname   string  `json:"hostname"`
				Interface  string  `json:"interface"`
				Speed      float64 `json:"speed"`
				Wifi       struct {
					Band string `json:"band"`
					Ssid string `json:"ssid"`
					Rssi int    `json:"rssi"`
				} `json:"wifi"`
				Ipv4 struct {
					Address            string `json:"address"`
					LeaseTimeRemaining int    `json:"leaseTimeRemaining"`
				} `json:"ipv4"`
				Ipv6 struct {
					LinkLocalAddress   string `json:"linkLocalAddress"`
					GlobalAddress      string `json:"globalAddress"`
					LeaseTimeRemaining int    `json:"leaseTimeRemaining"`
				} `json:"ipv6"`
			} `json:"config"`
		} `json:"hosts"`
	} `json:"hosts"`
}

type NetworkClient struct {
	httpClient *client.Client
}

func NewSagemcomNetworkClient(client *client.Client) *NetworkClient {
	return &NetworkClient{
		httpClient: client,
	}
}

func (c *NetworkClient) GetHosts(connectedOnly bool) (*NetworkHostResponse, error) {
	resp, err := c.httpClient.GET(fmt.Sprintf("network/hosts?connectedOnly=%t", connectedOnly), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get ipv4 dhcp: %w", err)
	}

	defer resp.Body.Close()

	var response NetworkHostResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}
