package models

type Answer struct {
	Id         int    `json:"id" gorm:"column:Id;primaryKey;auto_increment;not_null"`
	UserId     int    `json:"userid" gorm:"column:UserId;not_null"`
	PostId     int    `json:"postid" gorm:"column:PostId;not_null"`
	Post       Post   `gorm:"foreignkey:PostId;references:Id"`
	User       User   `json:"User" gorm:"foreignkey:UserId;references:Id"`
	Answer     string `json:"answer" gorm:"column:Answer"`
	PostDate   string `json:"postdate" gorm:"column:PostDate"`
	Likes      int    `json:"likes" gorm:"column:Likes"`
	Dislikes   int    `json:"dislikes" gorm:"column:Dislikes"`
	Is_Deleted int    `json:"is_deleted" gorm:"column:Is_Deleted"`
}

func (Answer) TableName() string {
	return "answer"
}
