package entity

type UserPost struct {
	Id     int `json:"id"`
	UserId int `json:"userid"`
	PostId int `json:"postid"`
	Grade  int `json:"grade"`
}

func (UserPost) TableName() string {
	return "userpost"
}
