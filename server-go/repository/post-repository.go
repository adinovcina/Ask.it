package repository

import (
	"time"

	"github.com/adinovcina/entity"
	"gorm.io/gorm"
)

type PostRepository interface {
	GetAll() []entity.Post
	Insert(entity.Post) entity.Post
	MostLikedPost() []entity.MostLikedPost
	MyPosts(int) []entity.Post
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
	var post entity.Post
	db.connection.Last(&post)
	return post
}

func (db *postConnection) MostLikedPost() []entity.MostLikedPost {
	var ml []entity.MostLikedPost
	db.connection.Table("userpost").Preload("Post").Model(&entity.MostAnswers{}).
		Select("count(postid) as Likes,post.title,post.PostDate").Joins("LEFT JOIN post on post.id = userpost.postid").
		Where("grade = 1").Group("postid").Find(&ml)
	return ml
}

func (db *postConnection) MyPosts(userId int) []entity.Post {
	var posts []entity.Post
	db.connection.Preload("User").Order("postdate desc").Where("userid = ?", userId).Limit(20).Find(&posts)
	return posts
}
