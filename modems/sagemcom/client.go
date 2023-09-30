package sagemcom

import (
	"fmt"
	"github.com/peacecwz/modem-sdk/pkg/client"
)

type Client struct {
	httpClient     *client.Client
	PortForwarding *PortForwardingClient
	User           *UserClient
	LoggedSession  *LoginResponse
}

func NewSagemcomClient(baseAddress string) *Client {
	httpClient := client.NewClient(baseAddress, map[string]string{})
	return &Client{
		httpClient:     httpClient,
		PortForwarding: NewSagemcomPortForwardingClient(httpClient),
		User:           NewSagemcomUserClient(httpClient),
	}
}

func (c *Client) Login(password string) (bool, error) {
	resp, err := c.User.Login(password)
	if err != nil {
		return false, fmt.Errorf("failed to login: %w", err)
	}

	c.httpClient.AddHeader("Authorization", fmt.Sprintf("Bearer %s", resp.Created.Token))
	c.LoggedSession = resp

	return true, nil
}

func (c *Client) Logout() (bool, error) {
	isSuccess, err := c.User.Logout(c.LoggedSession.Created.UserID, c.LoggedSession.Created.Token)
	if err != nil {
		return false, fmt.Errorf("failed to login: %w", err)
	}

	if isSuccess {
		c.httpClient.RemoveHeader("Authorization")
		c.LoggedSession = nil
	}

	return isSuccess, nil
}

func (c *Client) IsLoggedIn() bool {
	return c.LoggedSession != nil
}
