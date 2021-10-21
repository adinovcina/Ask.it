package repository

import (
	"time"

	"github.com/adinovcina/entity"
	"gorm.io/gorm"
)

type PostRepository interface {
	GetAll() []entity.Post
	Insert(entity.Post) entity.Post
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
