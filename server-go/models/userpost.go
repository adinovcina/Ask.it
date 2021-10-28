package models

type UserPost struct {
	Id     int  `json:"id" gorm:"column:Id;primaryKey;auto_increment;not_null"`
	UserId int  `json:"userid" gorm:"column:UserId;not_null"`
	PostId int  `json:"postid" gorm:"column:PostId;not_null"`
	User   User `gorm:"foreignkey:UserId;references:Id"`
	Post   Post `gorm:"foreignkey:PostId;references:Id"`
	Grade  int  `json:"grade" gorm:"column:Grade"`
}

func (UserPost) TableName() string {
	return "userpost"
}
