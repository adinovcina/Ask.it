package service

import (
	"github.com/adinovcina/entity"
	"github.com/adinovcina/repository"
)

type AnswerService interface {
	GetAll() []entity.Answer
	Insert(entity.Answer) entity.Answer
	// GetAllAnswersByPostId(id int) []models.Answer
	// // UpdateAnswerMark(models.Answer) models.Answer
	// VerifyIfAnswerMarkExist(models.Answer) models.Answer
	Update(entity.Answer)
	UpdateGrade(string, int) []entity.Answer
	// MostAnswers() []models.MostAnswers
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

// func (service *answerService) GetAllAnswersByPostId(id int) []models.Answer {
// 	res := service.answerRepository.GetAllAnswersByPostId(id)
// 	return res
// }

// // func (service *answerService) UpdateAnswerMark(answ models.Answer) models.Answer {
// // 	res := service.answerRepository.UpdateAnswerMark(answ)
// // 	return res
// // }

// func (service *answerService) VerifyIfAnswerMarkExist(ans models.Answer) models.Answer {
// 	res := service.answerRepository.VerifyIfAnswerMarkExist(ans)
// 	return res
// }

func (service *answerService) Update(ans entity.Answer) {
	service.answerRepository.Update(ans)
}

func (service *answerService) UpdateGrade(str string, postId int) []entity.Answer {
	res := service.answerRepository.UpdateGrade(str, postId)
	return res
}

// func (service *answerService) MostAnswers() []models.MostAnswers {
// 	res := service.answerRepository.MostAnswers()
// 	return res
// }
