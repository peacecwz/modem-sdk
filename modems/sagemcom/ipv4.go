package sagemcom

import (
	"encoding/json"
	"fmt"
	"github.com/peacecwz/modem-sdk/pkg/client"
)

type IPV4InfoResponse struct {
	Info struct {
		LanIPAddress  string `json:"lanIpAddress"`
		LanSubnet     string `json:"lanSubnet"`
		LanSubnetMask string `json:"lanSubnetMask"`
	} `json:"info"`
}

type IPV4DHCPResponse struct {
	DHCP struct {
		Enable     bool   `json:"enable"`
		MinAddress string `json:"minAddress"`
		MaxAddress string `json:"maxAddress"`
		SubnetMask string `json:"subnetMask"`
		LeaseTime  int    `json:"leaseTime"`
		MaxIps     int    `json:"maxIps"`
		AllowedIps []struct {
			LanAllowedSubnetIp   string `json:"lanAllowedSubnetIp"`
			LanAllowedSubnetMask string `json:"lanAllowedSubnetMask"`
		} `json:"allowedIps"`
		BlockedIps []struct {
			LanBlockedSubnetIp   string `json:"lanBlockedSubnetIp"`
			LanBlockedSubnetMask string `json:"lanBlockedSubnetMask"`
		} `json:"blockedIps"`
	} `json:"dhcp"`
}

type IPV4Client struct {
	httpClient *client.Client
}

func NewSagemcomIPV4Client(client *client.Client) *IPV4Client {
	return &IPV4Client{
		httpClient: client,
	}
}

func (c *IPV4Client) GetInfo() (*IPV4InfoResponse, error) {
	resp, err := c.httpClient.GET("network/ipv4/info", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get ipv4 info: %w", err)
	}

	defer resp.Body.Close()

	var response IPV4InfoResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}

func (c *IPV4Client) GetDHCP() (*IPV4DHCPResponse, error) {
	resp, err := c.httpClient.GET("network/ipv4/dhcp", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get ipv4 dhcp: %w", err)
	}

	defer resp.Body.Close()

	var response IPV4DHCPResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}
