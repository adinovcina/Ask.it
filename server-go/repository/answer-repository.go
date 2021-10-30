package repository

import (
	"time"

	"github.com/adinovcina/entity"
	"github.com/pusher/pusher-http-go"
	"gorm.io/gorm"
)

type AnswerRepository interface {
	GetAll() []entity.Answer
	Insert(entity.Answer) entity.Answer
	MostAnswers() []entity.MostAnswers
	EditAnswer(entity.Answer) entity.Answer
	DeleteAnswer(int) entity.Answer
	SendNotification(entity.Answer)
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
	db.connection.Preload("User").Last(&ans)
	db.SendNotification(ans)
	return ans
}

func (db *answerConnection) SendNotification(ans entity.Answer) {
	var userWhoPosted entity.Post
	db.connection.Where("id = ?", ans.PostId).First(&userWhoPosted)
	if userWhoPosted.UserId != ans.User.Id {
		notif := entity.Notification{PostId: ans.PostId,
			User_sending: ans.User.FirstName + " " + ans.User.LastName, User_receivingID: userWhoPosted.UserId}
		pusherClient := pusher.Client{
			AppID:   "1289244",
			Key:     "d88f98f756818528eb43",
			Secret:  "6049b8b0725a1ea31c4a",
			Cluster: "eu",
			Secure:  true,
		}
		db.connection.Exec(`INSERT INTO notification (User_sending, User_receiving, PostId, CommentDate) 
	 VALUES (?, ?, ?, ?)`, ans.User.FirstName+" "+ans.User.LastName,
			userWhoPosted.UserId, ans.PostId, ans.PostDate)
		pusherClient.Trigger("notification", "sendNotif", notif)
	}
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
