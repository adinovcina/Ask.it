package service

import (
	"github.com/adinovcina/entity"
	"github.com/adinovcina/repository"
)

type PostService interface {
	GetAll() []entity.Post
	Insert(entity.Post) entity.Post
	Update(entity.Post)
	UpdateGrade(string, int) entity.Post
	MostLikedPost() []entity.MostLikedPost
	MyPosts(int) []entity.Post
}

type postService struct {
	postRepository repository.PostRepository
}

func NewPostService(postRep repository.PostRepository) PostService {
	return &postService{
		postRepository: postRep,
	}
}

func (service *postService) GetAll() []entity.Post {
	res := service.postRepository.GetAll()
	return res
}

func (service *postService) Insert(post entity.Post) entity.Post {
	res := service.postRepository.Insert(post)
	return res
}

func (service *postService) Update(post entity.Post) {
	service.postRepository.Update(post)
}

func (service *postService) UpdateGrade(str string, postId int) entity.Post {
	res := service.postRepository.UpdateGrade(str, postId)
	return res
}

func (service *postService) MostLikedPost() []entity.MostLikedPost {
	res := service.postRepository.MostLikedPost()
	return res
}

func (service *postService) MyPosts(userId int) []entity.Post {
	res := service.postRepository.MyPosts(userId)
	return res
}
