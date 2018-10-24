package instagram

type Media struct {
	ID          string    `json:"id"`
	User        *User     `json:"user"`
	Comments    *Comments `json:"comments"`
	Caption     *Caption  `json:"caption"`
	Likes       *Likes    `json:"likes"`
	Type        MediaType `json:"type"`
	Images      *Images   `json:"images"`
	Videos      *Videos   `json:"videos"`
	Filter      string    `json:"filter"`
	Tags        []string  `json:"tags"`
	Location    *Location `json:"location"`
	Link        string    `json:"string"`
	CreatedTime int64     `json:"created_time"`
}

type RecentMediaResponse struct {
	Data []*Media `json:"data"`
}
