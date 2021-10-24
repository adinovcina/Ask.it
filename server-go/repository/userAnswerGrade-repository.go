package repository

import (
	"github.com/adinovcina/entity"
	"gorm.io/gorm"
)

type AnswerPostRepository interface {
	GetAll() []entity.AnswerPost
	Insert(entity.AnswerPost) entity.AnswerPost
	UpdateAnswerMark(entity.AnswerPost) entity.AnswerPost
	VerifyIfGradeExist(entity.AnswerPost) entity.AnswerPost
	Verify(entity.AnswerPost) bool
}

type answerpostConnection struct {
	connection *gorm.DB
}

func NewAnswerPostRepository(db *gorm.DB) AnswerPostRepository {
	return &answerpostConnection{
		connection: db,
	}
}

func (db *answerpostConnection) GetAll() []entity.AnswerPost {
	var answerPost []entity.AnswerPost
	db.connection.Find(&answerPost)
	return answerPost
}

func (db *answerpostConnection) Insert(newAnswerPost entity.AnswerPost) entity.AnswerPost {
	db.connection.Exec(`INSERT INTO useranswer (answerid,userid,postid,grade) VALUES (?, ?, ?, ?)`, newAnswerPost.AnswerId, newAnswerPost.UserId, newAnswerPost.PostId, newAnswerPost.Grade)
	return newAnswerPost
}

func (db *answerpostConnection) UpdateAnswerMark(newAnswerPost entity.AnswerPost) entity.AnswerPost {
	db.connection.Model(entity.AnswerPost{}).Where("userid = ? and postid = ? and answerid = ?", newAnswerPost.UserId,
		newAnswerPost.PostId, newAnswerPost.AnswerId).Updates(entity.UserPost{Grade: newAnswerPost.Grade})
	return newAnswerPost
}

func (db *answerpostConnection) VerifyIfGradeExist(newAnswer entity.AnswerPost) entity.AnswerPost {
	var answer entity.AnswerPost
	db.connection.Where("answerid = ? and userid = ? AND postid = ? AND grade != ?", newAnswer.AnswerId,
		newAnswer.UserId, newAnswer.PostId, newAnswer.Grade).First(&answer)
	return answer
}

func (db *answerpostConnection) Verify(newAnswer entity.AnswerPost) bool {
	var answer entity.AnswerPost
	db.connection.Where("answerid = ? and userid = ? AND postid = ? AND grade = ?",
		newAnswer.AnswerId, newAnswer.UserId, newAnswer.PostId, newAnswer.Grade).First(&answer)
	return answer.Id != 0
}
