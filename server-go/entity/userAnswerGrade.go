package entity

type AnswerPost struct {
	Id       int `json:"id"`
	UserId   int `json:"userid"`
	AnswerId int `json:"answerid"`
	PostId   int `json:"postid"`
	Grade    int `json:"grade"`
}

func (AnswerPost) TableName() string {
	return "useranswer"
}
