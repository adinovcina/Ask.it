package repository

import (
	"github.com/adinovcina/entity"
	"github.com/adinovcina/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(entity.User) entity.User
	UpdateUser(entity.User) entity.User
	VerifyCredential(string, string) interface{}
	IsDuplicatedEmail(string) bool
	FindUser(string) entity.User
	EditOldPassword(models.UserProfile)
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func (db *userConnection) InsertUser(user entity.User) entity.User {
	password := HashPassword(user.Password)
	db.connection.Exec(`INSERT INTO user (FirstName, LastName, Email, PasswordHash) VALUES (?, ?, ?, ?)`,
		user.FirstName, user.LastName, user.Email, password)
	return user
}

func (db *userConnection) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		password := HashPassword(user.Password)
		user.Password = password
		db.connection.Save(&user)
		return user
	}
	return entity.User{}
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).Find(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicatedEmail(email string) bool {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user)
	return user.Id == 0
}

func (db *userConnection) FindUser(email string) entity.User {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnection) EditOldPassword(user models.UserProfile) {
	hash := HashPassword(user.NewPassword)

	var u entity.User
	db.connection.Where("email = ?", user.Email).First(&u)

	db.connection.Model(&u).Updates(map[string]interface{}{"firstname": user.FirstName,
		"lastname": user.LastName, "passwordhash": hash})
}
