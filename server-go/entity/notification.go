package entity

type Notification struct {
	Id               int    `json:"id"`
	User_sending     string `json:"User_sending"`
	User_receivingID int    `json:"User_receiving"`
	PostId           int    `json:"postid"`
	Post             Post   `json:"Post"`
	CommentDate      string `json:"commentdate"`
}

func (Notification) TableName() string {
	return "notification"
}
