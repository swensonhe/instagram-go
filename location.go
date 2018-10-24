package instagram

type Location struct {
	ID            string  `json:"id"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	StreetAddress string  `json:"street_address"`
	Name          string  `json:"name"`
}
