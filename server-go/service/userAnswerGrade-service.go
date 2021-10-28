package service

import (
	"github.com/adinovcina/entity"
	"github.com/adinovcina/repository"
)

type AnswerPostService interface {
	GetAll() []entity.AnswerPost
	Insert(entity.AnswerPost) entity.AnswerPost
	UpdateAnswerMark(entity.AnswerPost) entity.AnswerPost
	VerifyIfGradeExist(entity.AnswerPost) entity.AnswerPost
	Verify(entity.AnswerPost) bool
	VerifyIfDataExist(entity.AnswerPost) bool
}

type answerpostService struct {
	answerpostRepository repository.AnswerPostRepository
}

func NewAnswerPostService(answerRep repository.AnswerPostRepository) AnswerPostService {
	return &answerpostService{
		answerpostRepository: answerRep,
	}
}

func (service *answerpostService) GetAll() []entity.AnswerPost {
	res := service.answerpostRepository.GetAll()
	return res
}

func (service *answerpostService) Insert(answerPost entity.AnswerPost) entity.AnswerPost {
	res := service.answerpostRepository.Insert(answerPost)
	return res
}

func (service *answerpostService) UpdateAnswerMark(answerPost entity.AnswerPost) entity.AnswerPost {
	res := service.answerpostRepository.UpdateAnswerMark(answerPost)
	return res
}

func (service *answerpostService) VerifyIfGradeExist(answerPost entity.AnswerPost) entity.AnswerPost {
	res := service.answerpostRepository.VerifyIfGradeExist(answerPost)
	return res
}

func (service *answerpostService) Verify(answerPost entity.AnswerPost) bool {
	res := service.answerpostRepository.Verify(answerPost)
	return res
}

func (service *answerpostService) VerifyIfDataExist(answerPost entity.AnswerPost) bool {
	res := service.answerpostRepository.VerifyIfDataExist(answerPost)
	return res
}
