package entity

type Post struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	PostDate string `json:"postdate"`
	User     User   `json:"User"`
	UserId   int    `json:"Userid"`
}

func (Post) TableName() string {
	return "post"
}
