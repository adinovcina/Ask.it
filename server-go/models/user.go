package models

type User struct {
	Id           int    `json:"Id" gorm:"column:Id;primaryKey;auto_increment;not_null"`
	Email        string `json:"Email" gorm:"column:Email"`
	FirstName    string `json:"FirstName" gorm:"column:FirstName"`
	LastName     string `json:"LastName" gorm:"column:LastName"`
	Passwordhash string `json:"Passwordhash" gorm:"column:Passwordhash"`
}

func (User) TableName() string {
	return "user"
}
