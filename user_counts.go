package instagram

type UserCounts struct {
	Media int64 `json:"media"`
	Follows int64 `json:"follows"`
	FollowedBy int64 `json:"followed_by"`
}
