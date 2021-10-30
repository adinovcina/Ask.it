package repository

import (
	"github.com/adinovcina/entity"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	GetAll(int) []entity.Notification
}

type notificationConnection struct {
	connection *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationConnection{
		connection: db,
	}
}

func (db *notificationConnection) GetAll(userId int) []entity.Notification {
	var notifications []entity.Notification
	db.connection.Preload("Post").Where("user_receiving = ?", userId).Find(&notifications)
	return notifications
}
