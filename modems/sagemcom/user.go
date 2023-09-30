package sagemcom

import (
	"encoding/json"
	"fmt"
	"github.com/peacecwz/modem-sdk/pkg/client"
)

type LoginRequest struct {
	Password string `json:"password"`
}

type LoginResponse struct {
	Created struct {
		Token     string `json:"token"`
		UserLevel string `json:"userLevel"`
		UserID    int    `json:"userId"`
	} `json:"created"`
}

type UserClient struct {
	httpClient *client.Client
}

func NewSagemcomUserClient(client *client.Client) *UserClient {
	return &UserClient{
		httpClient: client,
	}
}

func (c *UserClient) Login(password string) (*LoginResponse, error) {
	reqBody := LoginRequest{
		Password: password,
	}

	resp, err := c.httpClient.POST("user/login", nil, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to login: %w, resp: %+v", err, resp)
	}

	defer resp.Body.Close()

	var loginResponse LoginResponse
	err = json.NewDecoder(resp.Body).Decode(&loginResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to decode login response: %w", err)
	}

	return &loginResponse, nil
}

func (c *UserClient) Logout(userId int, token string) (bool, error) {
	resp, err := c.httpClient.DELETE(fmt.Sprintf("user/%d/token/%s", userId, token), nil)
	if err != nil {
		return false, fmt.Errorf("failed to logout: %w, resp: %+v", err, resp)
	}

	defer resp.Body.Close()

	return resp.StatusCode == 204, nil
}
