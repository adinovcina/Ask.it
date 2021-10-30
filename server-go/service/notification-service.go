package service

import (
	"github.com/adinovcina/entity"
	"github.com/adinovcina/repository"
)

type NotificationService interface {
	GetAll(int) []entity.Notification
}

type notificationService struct {
	notificationRepository repository.NotificationRepository
}

func NewNotificationService(notification repository.NotificationRepository) NotificationService {
	return &notificationService{
		notificationRepository: notification,
	}
}

func (service *notificationService) GetAll(userId int) []entity.Notification {
	res := service.notificationRepository.GetAll(userId)
	return res
}
