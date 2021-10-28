package models

type UserProfile struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Passwordhash string `json:"passwordhash"`
	OldPassword  string `json:"oldpassword"`
	NewPassword  string `json:"newpassword"`
}

func (UserProfile) TableName() string {
	return "user"
}
