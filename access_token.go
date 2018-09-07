package instagram

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	User User `json:"user"`
}
