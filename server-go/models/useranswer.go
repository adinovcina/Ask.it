package models

type AnswerPost struct {
	Id       int    `json:"id" gorm:"column:Id;primaryKey;auto_increment;not_null"`
	UserId   int    `json:"userid" gorm:"column:UserId;not_null"`
	User     User   `gorm:"foreignkey:UserId;references:Id"`
	AnswerId int    `json:"answerid" gorm:"column:AnswerId;not_null"`
	Answer   Answer `gorm:"foreignkey:AnswerId;references:Id"`
	Post     Post   `gorm:"foreignkey:PostId;references:Id"`
	PostId   int    `json:"postid" gorm:"column:PostId;not_null"`
	Grade    int    `json:"grade" gorm:"column:Grade"`
}

func (AnswerPost) TableName() string {
	return "useranswer"
}
