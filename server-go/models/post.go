package models

type Post struct {
	Id       int    `json:"id" gorm:"column:Id;primaryKey;auto_increment;not_null"`
	Title    string `json:"title" gorm:"column:Title"`
	PostDate string `json:"postdate" gorm:"column:PostDate"`
	UserId   int    `json:"Userid" gorm:"column:UserId;not_null"`
	User     User   `json:"User" gorm:"foreignkey:UserId;references:Id"`
	Likes    int    `json:"likes" gorm:"column:Likes"`
	Dislikes int    `json:"dislikes" gorm:"column:Dislikes"`
}

func (Post) TableName() string {
	return "post"
}
