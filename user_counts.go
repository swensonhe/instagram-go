package instagram

type UserCounts struct {
	Media int `json:"media"`
	Follows int `json:"follows"`
	FollowedBy int `json:"followed_by"`
}
