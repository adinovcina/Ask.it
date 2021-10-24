package repository

import (
	"time"

	"github.com/adinovcina/entity"
	"gorm.io/gorm"
)

type PostRepository interface {
	GetAll() []entity.Post
	Insert(entity.Post) entity.Post
	Update(entity.Post)
	UpdateGrade(string, int) []entity.Post
}

type postConnection struct {
	connection *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postConnection{
		connection: db,
	}
}

func (db *postConnection) GetAll() []entity.Post {
	var posts []entity.Post
	db.connection.Preload("User").Find(&posts)
	return posts
}

func (db *postConnection) Insert(newPost entity.Post) entity.Post {
	now := time.Now()
	formatedDate := now.Format("2006-01-02 15:04:05")
	newPost.PostDate = formatedDate
	db.connection.Exec(`INSERT INTO post (title, postdate, likes, dislikes, userid) VALUES (?, ?, ?, ?, ?)`,
		newPost.Title, newPost.PostDate, 0, 0, newPost.UserId)
	return newPost
}

func (db *postConnection) Update(newPost entity.Post) {
	if newPost.Likes != 0 {
		db.connection.Model(entity.Post{}).Where("id = ?", newPost.Id).
			UpdateColumn("Likes", gorm.Expr("Likes + ?", 1))
	} else {
		db.connection.Model(entity.Post{}).Where("id = ?", newPost.Id).
			UpdateColumn("Dislikes", gorm.Expr("Dislikes + ?", 1))
	}
}

func (db *postConnection) UpdateGrade(str string, postId int) []entity.Post {
	if str == "dislike" {
		db.connection.Model(entity.Post{}).Where("id = ?", postId).
			UpdateColumn("Dislikes", gorm.Expr("Dislikes - ?", 1))
	} else {
		db.connection.Model(entity.Post{}).Where("id = ?", postId).
			UpdateColumn("Likes", gorm.Expr("Likes - ?", 1))
	}
	return db.GetAll()
}
