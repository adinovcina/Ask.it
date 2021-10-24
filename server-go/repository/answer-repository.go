package repository

import (
	"time"

	"github.com/adinovcina/entity"
	"gorm.io/gorm"
)

type AnswerRepository interface {
	GetAll() []entity.Answer
	Insert(entity.Answer) entity.Answer
	// GetAllAnswersByPostId(id int) []models.Answer
	// // UpdateAnswerMark(models.Answer) models.Answer
	// VerifyIfAnswerMarkExist(models.Answer) models.Answer
	Update(entity.Answer)
	UpdateGrade(string, int) []entity.Answer
	// MostAnswers() []models.MostAnswers
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
	db.connection.Preload("User").Find(&answers)
	return answers
}

func (db *answerConnection) Insert(newAnswer entity.Answer) entity.Answer {
	now := time.Now()
	formatedDate := now.Format("2006-01-02 15:04:05")
	newAnswer.PostDate = formatedDate
	db.connection.Exec(`INSERT INTO answer (userid,postid,answer,postdate,likes,dislikes)
	VALUES (?, ?, ?, ?, ?, ?)`,
		newAnswer.UserId, newAnswer.PostId, newAnswer.Answer, formatedDate, 0, 0)
	return newAnswer
}

// func (db *answerConnection) GetAllAnswersByPostId(id int) []models.Answer {
// 	var answers []models.Answer
// 	db.connection.Where("postid = ?", id).Find(&answers)
// 	return answers
// }

// // func (db *answerConnection) UpdateAnswerMark(newAnswer models.Answer) models.Answer {
// // 	var answer models.Answer
// // 	db.connection.Where("id = ?", newAnswer).Update("answermark", newAnswer.AnswerMark)
// // 	return answer
// // }

// func (db *answerConnection) VerifyIfAnswerMarkExist(newAnswer models.Answer) models.Answer {
// 	var answer models.Answer
// 	db.connection.Where("userid = ? AND postid = ?", newAnswer.UserId, newAnswer.PostId).First(&answer)
// 	return answer
// }

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

// func (db *answerConnection) MostAnswers() []models.MostAnswers {
// 	var mostAnsw []models.MostAnswers
// 	db.connection.Preload("User").Model(&models.MostAnswers{}).Select("count(userid) as NumberOfAnswers, userid as UserId").
// 		Group("userid").Find(&mostAnsw)
// 	return mostAnsw
// }
