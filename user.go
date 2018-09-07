package instagram

type User struct {
	Id string `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
	Bio string `json:"bio"`
	Website string `json:"website"`
	IsBusiness bool `json:"is_business"`
	Counts UserCounts `json:"counts"`
}

type UserResponse struct {
	Data User `json:"data"`
}