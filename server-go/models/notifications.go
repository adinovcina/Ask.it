package models

type Notifications struct {
	Id               int  `json:"id" gorm:"column:Id;primaryKey;auto_increment;not_null"`
	User_sendingId   int  `json:"User_sending" gorm:"column:User_sending;not_null"`
	User_receivingId int  `json:"User_receiving" gorm:"column:User_receiving;not_null"`
	UserSending      User `json:"UserSending" gorm:"foreignkey:User_sendingId;references:Id"`
	UserReceiving    User `json:"UserReceiving" gorm:"foreignkey:User_receivingId;references:Id"`
	PostId           int  `json:"postid" gorm:"column:PostId;not_null"`
	Post             Post `json:"Post" gorm:"foreignkey:PostId;references:Id"`
	Seen_by_user     bool `json:"Seen_by_user" gorm:"column:Seen_by_user"`
}

func (Notifications) TableName() string {
	return "notifications"
}
