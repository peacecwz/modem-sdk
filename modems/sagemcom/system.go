package sagemcom

import (
	"encoding/json"
	"fmt"
	"github.com/peacecwz/modem-sdk/pkg/client"
)

type SystemInfoResponse struct {
	Info struct {
		ModelName       string `json:"modelName"`
		SoftwareVersion string `json:"softwareVersion"`
		HardwareVersion string `json:"hardwareVersion"`
	} `json:"info"`
}

type SystemGatewayProvisioningResponse struct {
	Provisioning struct {
		Mode       string `json:"mode"`
		MacAddress string `json:"macAddress"`
		IPv4       struct {
			Address        string   `json:"address"`
			DefaultGateway string   `json:"defaultGateway"`
			LeaseTime      int      `json:"leaseTime"`
			ExpireTime     int      `json:"expireTime"`
			DNSServers     []string `json:"dnsServers"`
		} `json:"ipv4"`
		IPv6 struct {
			GlobalAddress  string   `json:"globalAddress"`
			LinkAddress    string   `json:"linkAddress"`
			DefaultGateway string   `json:"defaultGateway"`
			LeaseTime      int      `json:"leaseTime"`
			ExpireTime     int      `json:"expireTime"`
			DNSServers     []string `json:"dnsServers"`
		} `json:"ipv6"`
		DsLite struct {
			Enable bool `json:"enable"`
		} `json:"dsLite"`
	} `json:"provisioning"`
}

type SystemClient struct {
	httpClient *client.Client
}

func NewSagemcomSystemClient(client *client.Client) *SystemClient {
	return &SystemClient{
		httpClient: client,
	}
}

func (c *SystemClient) GetSystemInfo() (*SystemInfoResponse, error) {
	resp, err := c.httpClient.GET("system/info", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get system info: %w", err)
	}

	defer resp.Body.Close()

	var response SystemInfoResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}

func (c *SystemClient) GetSystemGatewayInfoProvisioning() (*SystemGatewayProvisioningResponse, error) {
	resp, err := c.httpClient.GET("system/gateway/provisioning", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get system gateway info: %w", err)
	}

	defer resp.Body.Close()

	var response SystemGatewayProvisioningResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}
