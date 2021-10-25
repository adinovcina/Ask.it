package repository

import (
	"time"

	"github.com/adinovcina/entity"
	"gorm.io/gorm"
)

type AnswerRepository interface {
	GetAll() []entity.Answer
	Insert(entity.Answer) entity.Answer
	Update(entity.Answer)
	UpdateGrade(string, int) []entity.Answer
	MostAnswers() []entity.MostAnswers
	EditAnswer(entity.Answer) entity.Answer
	DeleteAnswer(int) entity.Answer
}

type answerConnection struct {
	connection *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) AnswerRepository {
	return &answerConnection{
		connection: db,
	}
}

func (db *answerConnection) GetAll() []entity.Answer {
	var answers []entity.Answer
	db.connection.Preload("User").Where("is_deleted = 0").Find(&answers)
	return answers
}

func (db *answerConnection) Insert(newAnswer entity.Answer) entity.Answer {
	now := time.Now()
	formatedDate := now.Format("2006-01-02 15:04:05")
	newAnswer.PostDate = formatedDate
	db.connection.Exec(`INSERT INTO answer (userid,postid,answer,postdate,likes,dislikes,is_deleted)
	VALUES (?, ?, ?, ?, ?, ?, ?)`,
		newAnswer.UserId, newAnswer.PostId, newAnswer.Answer, formatedDate, 0, 0, 0)
	var ans entity.Answer
	db.connection.Last(&ans)
	return ans
}

func (db *answerConnection) Update(newPost entity.Answer) {
	if newPost.Likes != 0 {
		db.connection.Model(entity.Answer{}).Where("id = ?", newPost.Id).
			UpdateColumn("Likes", gorm.Expr("Likes + ?", 1))
	} else {
		db.connection.Model(entity.Answer{}).Where("id = ?", newPost.Id).
			UpdateColumn("Dislikes", gorm.Expr("Dislikes + ?", 1))
	}
}

func (db *answerConnection) UpdateGrade(str string, postId int) []entity.Answer {
	if str == "dislike" {
		db.connection.Model(entity.Answer{}).Where("id = ?", postId).
			UpdateColumn("Dislikes", gorm.Expr("Dislikes - ?", 1))
	} else {
		db.connection.Model(entity.Answer{}).Where("id = ?", postId).
			UpdateColumn("Likes", gorm.Expr("Likes - ?", 1))
	}
	return db.GetAll()
}

func (db *answerConnection) MostAnswers() []entity.MostAnswers {
	var mostAnsw []entity.MostAnswers
	db.connection.Preload("User").Model(&entity.MostAnswers{}).Select("count(userid) as NumberOfAnswers, userid as UserId").
		Where("is_deleted = 0").Group("userid").Where("").Find(&mostAnsw)
	return mostAnsw
}

func (db *answerConnection) EditAnswer(answer entity.Answer) entity.Answer {
	db.connection.Model(entity.Answer{}).Where("id = ?", answer.Id).
		UpdateColumn("Answer", answer.Answer)
	var ans entity.Answer
	db.connection.Where("id = ?", answer.Id).First(&ans)
	return ans
}

func (db *answerConnection) DeleteAnswer(answerId int) entity.Answer {
	db.connection.Model(entity.Answer{}).Where("id = ?", answerId).
		UpdateColumn("Is_Deleted", 1)
	var ans entity.Answer
	db.connection.Where("id = ?", answerId).First(&ans)
	return ans
}
