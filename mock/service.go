package mock

import "github.com/swensonhe/instagram-go"

type APIClient struct {
	GetAccessTokenFn func(code string) (*instagram.AccessTokenResponse, error)
	GetAccessTokenInvoked bool

	GetSelfFn func(token string) (*instagram.UserResponse, error)
	GetSelfInvoked bool

	GetRecentMediaFn      func(token string, maxID string, minID string, count int) (*instagram.RecentMediaResponse, error)
	GetRecentMediaInvoked bool
}

func NewAPIClient() *APIClient {
	return &APIClient{
		GetAccessTokenFn: func(code string) (*instagram.AccessTokenResponse, error) {
			return &instagram.AccessTokenResponse{}, nil
		},
		GetSelfFn: func(token string) (*instagram.UserResponse, error) {
			return &instagram.UserResponse{}, nil
		},
		GetRecentMediaFn: func(token string, maxID string, minID string, count int) (*instagram.RecentMediaResponse, error) {
			return &instagram.RecentMediaResponse{}, nil
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

func (c *APIClient) GetRecentMedia(token string, maxID string, minID string, count int) (*instagram.RecentMediaResponse, error) {
	c.GetRecentMediaInvoked = true
	return c.GetRecentMediaFn(token, maxID, minID, count)
}
