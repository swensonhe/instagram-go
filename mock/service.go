package mock

import "github.com/swensonhe/instagram-go"

type APIClient struct {
	GetAccessTokenFn func(code string) (*instagram.AccessTokenResponse, error)
	GetAccessTokenInvoked bool

	GetSelfFn func(token string) (*instagram.UserResponse, error)
	GetSelfInvoked bool
}

func NewAPIClient() *APIClient {
	return &APIClient{
		GetAccessTokenFn: func(code string) (*instagram.AccessTokenResponse, error) {
			return &instagram.AccessTokenResponse{}, nil
		},
		GetSelfFn: func(token string) (*instagram.UserResponse, error) {
			return &instagram.UserResponse{}, nil
		},
	}
}

func (c *APIClient) GetAccessToken(code string) (*instagram.AccessTokenResponse, error) {
	c.GetAccessTokenInvoked = true
	return c.GetAccessTokenFn(code)
}

func (c *APIClient) GetSelf(token string) (*instagram.UserResponse, error) {
	c.GetSelfInvoked = true
	return c.GetSelfFn(token)
}
