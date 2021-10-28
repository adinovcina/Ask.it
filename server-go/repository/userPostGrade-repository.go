package repository

import (
	"github.com/adinovcina/entity"
	"gorm.io/gorm"
)

type UserPostRepository interface {
	GetAll() []entity.UserPost
	Insert(entity.UserPost) entity.UserPost
	UpdateAnswerMark(entity.UserPost) entity.UserPost
	VerifyIfGradeExist(entity.UserPost) entity.UserPost
	Verify(entity.UserPost) bool
	VerifyIfDataExist(entity.UserPost) bool
}

type userpostConnection struct {
	connection *gorm.DB
}

func NewUserPostRepository(db *gorm.DB) UserPostRepository {
	return &userpostConnection{
		connection: db,
	}
}

func (db *userpostConnection) GetAll() []entity.UserPost {
	var userposts []entity.UserPost
	db.connection.Find(&userposts)
	return userposts
}

func (db *userpostConnection) Insert(newUserPost entity.UserPost) entity.UserPost {
	db.connection.Exec(`INSERT INTO userpost (userid,postid,grade) VALUES (?, ?, ?)`, newUserPost.UserId, newUserPost.PostId, newUserPost.Grade)
	var userPost entity.UserPost
	db.connection.Where("userid = ? and postid = ?", newUserPost.UserId, newUserPost.PostId).First(&userPost)
	return userPost
}

func (db *userpostConnection) UpdateAnswerMark(newUserPost entity.UserPost) entity.UserPost {
	db.connection.Model(entity.UserPost{}).Where("userid = ? and postid = ?", newUserPost.UserId, newUserPost.PostId).
		Updates(entity.UserPost{Grade: newUserPost.Grade})
	var userPost entity.UserPost
	db.connection.Where("userid = ? and postid = ?", newUserPost.UserId, newUserPost.PostId).First(&userPost)
	return userPost
}

func (db *userpostConnection) VerifyIfGradeExist(newAnswer entity.UserPost) entity.UserPost {
	var answer entity.UserPost
	db.connection.Where("userid = ? AND postid = ? AND grade != ?", newAnswer.UserId, newAnswer.PostId,
		newAnswer.Grade).First(&answer)
	return answer
}

func (db *userpostConnection) Verify(newAnswer entity.UserPost) bool {
	var answer entity.UserPost
	db.connection.Where("userid = ? AND postid = ? AND grade = ?", newAnswer.UserId, newAnswer.PostId,
		newAnswer.Grade).First(&answer)
	return answer.Id != 0
}

func (db *userpostConnection) VerifyIfDataExist(newAnswer entity.UserPost) bool {
	var answer entity.UserPost
	db.connection.Where("userid = ? AND postid = ?", newAnswer.UserId, newAnswer.PostId).First(&answer)
	return answer.Id != 0
}
