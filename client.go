package instagram

import (
	"net/http"
	"fmt"
	"net/url"
	"io/ioutil"
	"encoding/json"
)

type APIClient interface {
	GetAccessToken(code string) (*AccessTokenResponse, error)
	GetSelf(token string) (*UserResponse, error)
	GetRecentMedia(token string, maxID string, minID string, count int) (*RecentMediaResponse, error)
}

type Client struct {
	*http.Client
	clientId string
	clientSecret string
	redirectUri string
	log func(arg interface{})
}

const (
	urlScheme = "https"
	urlHost = "api.instagram.com"

	formKeyClientId = "client_id"
	formKeyClientSecret = "client_secret"
	formKeyGrantType = "grant_type"
	formKeyRedirectUri = "redirect_uri"
	formKeyCode = "code"

	grantTypeAuthorizationCode = "authorization_code"
)

func NewClient(config *Config) *Client {
	return &Client{
		Client: &http.Client{},
		clientId: config.ClientId,
		clientSecret: config.ClientSecret,
		redirectUri: config.RedirectUri,
		log: func(arg interface{}) {
			if config.LoggingEnabled {
				fmt.Println(arg)
			}
		},
	}
}

// GetAccessToken returns an access token.
func (c *Client) GetAccessToken(code string) (*AccessTokenResponse, error) {
	u := url.URL{Scheme: urlScheme, Host: urlHost, Path: "/oauth/access_token"}

	form := url.Values{}
	form.Add(formKeyClientId, c.clientId)
	form.Add(formKeyClientSecret, c.clientSecret)
	form.Add(formKeyGrantType, grantTypeAuthorizationCode)
	form.Add(formKeyRedirectUri, c.redirectUri)
	form.Add(formKeyCode, code)

	res, err := c.Client.PostForm(u.String(), form)
	if err != nil {
		c.log(err)
		return nil, err
	}

	if res.StatusCode != 200 {
		respBody, _ := ioutil.ReadAll(res.Body)
		c.log(string(respBody))
		return nil, ErrInternal
	}

	var accessTokenResponse AccessTokenResponse
	err = bind(res, &accessTokenResponse)
	if err != nil {
		return nil, ErrInternal
	}

	return &accessTokenResponse, nil
}

// GetSelf returns a user.
func (c *Client) GetSelf(token string) (*UserResponse, error) {
	u := url.URL{Scheme: urlScheme, Host: urlHost, Path: "/v1/users/self/"}

	q := u.Query()
	q.Set("access_token", token)
	u.RawQuery = q.Encode()

	fmt.Println(u.String())

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		c.log(err)
	}

	res, err := c.Client.Do(req)
	if err != nil {
		c.log(err)
		return nil, err
	}

	if res.StatusCode != 200 {
		respBody, _ := ioutil.ReadAll(res.Body)
		c.log(string(respBody))
		return nil, ErrInternal
	}

	var resp UserResponse
	err = bind(res, &resp)
	if err != nil {
		return nil, ErrInternal
	}

	return &resp, nil
}

// GetRecentMedia returns a user's recent media.
func (c *Client) GetRecentMedia(token string, maxID string, minID string, count int) (*RecentMediaResponse, error) {
	u := url.URL{Scheme: urlScheme, Host: urlHost, Path: "/v1/users/self/media/recent"}

	q := u.Query()
	q.Set("access_token", token)

	if maxID != "" {
		q.Set("max_id", maxID)
	}

	if minID != "" {
		q.Set("min_id", minID)
	}

	if count != 0 {
		q.Set("count", fmt.Sprintf("%d", count))
	}

	u.RawQuery = q.Encode()

	fmt.Println(u.String())

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		c.log(err)
	}

	res, err := c.Client.Do(req)
	if err != nil {
		c.log(err)
		return nil, err
	}

	if res.StatusCode != 200 {
		respBody, _ := ioutil.ReadAll(res.Body)
		c.log(string(respBody))
		return nil, ErrInternal
	}

	var resp RecentMediaResponse
	err = bind(res, &resp)
	if err != nil {
		return nil, ErrInternal
	}

	return &resp, nil
}

func bind(res *http.Response, v interface{}) error {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}
