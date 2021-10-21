package service

import (
	"github.com/adinovcina/entity"
	"github.com/adinovcina/repository"
)

type PostService interface {
	GetAll() []entity.Post
	Insert(entity.Post) entity.Post
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
