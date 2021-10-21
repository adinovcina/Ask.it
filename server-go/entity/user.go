package entity

type User struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	Passwordhash string `json:"passwordhash"`
	Password     string `json:"password"`
	Token        string `json:"token"`
}

func (User) TableName() string {
	return "user"
}
