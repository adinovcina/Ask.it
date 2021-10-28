package service

import (
	"github.com/adinovcina/entity"
	"github.com/adinovcina/repository"
)

type UserPostService interface {
	GetAll() []entity.UserPost
	Insert(entity.UserPost) entity.UserPost
	UpdateAnswerMark(entity.UserPost) entity.UserPost
	VerifyIfGradeExist(entity.UserPost) entity.UserPost
	Verify(entity.UserPost) bool
	VerifyIfDataExist(entity.UserPost) bool
}

type userpostService struct {
	userpostRepository repository.UserPostRepository
}

func NewUserPostService(answerRep repository.UserPostRepository) UserPostService {
	return &userpostService{
		userpostRepository: answerRep,
	}
}

func (service *userpostService) GetAll() []entity.UserPost {
	res := service.userpostRepository.GetAll()
	return res
}

func (service *userpostService) Insert(userPost entity.UserPost) entity.UserPost {
	res := service.userpostRepository.Insert(userPost)
	return res
}

func (service *userpostService) UpdateAnswerMark(userPost entity.UserPost) entity.UserPost {
	res := service.userpostRepository.UpdateAnswerMark(userPost)
	return res
}

func (service *userpostService) VerifyIfGradeExist(userPost entity.UserPost) entity.UserPost {
	res := service.userpostRepository.VerifyIfGradeExist(userPost)
	return res
}

func (service *userpostService) Verify(userPost entity.UserPost) bool {
	res := service.userpostRepository.Verify(userPost)
	return res
}

func (service *userpostService) VerifyIfDataExist(userPost entity.UserPost) bool {
	res := service.userpostRepository.VerifyIfDataExist(userPost)
	return res
}
