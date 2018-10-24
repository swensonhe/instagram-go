package instagram

type Caption struct {
	ID          string `json:"id"`
	Text        string `json:"text"`
	From        *User  `json:"from"`
	CreatedTime int64  `json:"created_time"`
}
