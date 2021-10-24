package entity

type Answer struct {
	Id       int    `json:"id"`
	UserId   int    `json:"userid"`
	PostId   int    `json:"postid"`
	User     User   `json:"User"`
	Answer   string `json:"answer"`
	PostDate string `json:"postdate"`
	Likes    int    `json:"likes"`
	Dislikes int    `json:"dislikes"`
}

func (Answer) TableName() string {
	return "answer"
}
