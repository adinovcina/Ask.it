package service

import (
	"github.com/adinovcina/entity"
	"github.com/adinovcina/repository"
)

type AnswerService interface {
	GetAll() []entity.Answer
	Insert(entity.Answer) entity.Answer
	Update(entity.Answer)
	UpdateGrade(string, int) entity.Answer
	MostAnswers() []entity.MostAnswers
	EditAnswer(entity.Answer) entity.Answer
	DeleteAnswer(int) entity.Answer
}

type answerService struct {
	answerRepository repository.AnswerRepository
}

func NewAnswerService(answerRep repository.AnswerRepository) AnswerService {
	return &answerService{
		answerRepository: answerRep,
	}
}

func (service *answerService) GetAll() []entity.Answer {
	res := service.answerRepository.GetAll()
	return res
}

func (service *answerService) Insert(answer entity.Answer) entity.Answer {
	res := service.answerRepository.Insert(answer)
	return res
}

func (service *answerService) Update(ans entity.Answer) {
	service.answerRepository.Update(ans)
}

func (service *answerService) UpdateGrade(str string, postId int) entity.Answer {
	res := service.answerRepository.UpdateGrade(str, postId)
	return res
}

func (service *answerService) MostAnswers() []entity.MostAnswers {
	res := service.answerRepository.MostAnswers()
	return res
}

func (service *answerService) EditAnswer(ans entity.Answer) entity.Answer {
	res := service.answerRepository.EditAnswer(ans)
	return res
}

func (service *answerService) DeleteAnswer(ans int) entity.Answer {
	res := service.answerRepository.DeleteAnswer(ans)
	return res
}
