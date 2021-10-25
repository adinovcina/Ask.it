package entity

type MostAnswers struct {
	UserId          int  `json:"userid"`
	User            User `json:"user"`
	NumberOfAnswers int  `json:"numberofanswers"`
}

func (MostAnswers) TableName() string {
	return "answer"
}
