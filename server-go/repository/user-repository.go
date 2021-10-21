package repository

import (
	"github.com/adinovcina/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(entity.User) entity.User
	UpdateUser(entity.User) entity.User
	CheckIfEmailExist(string) bool
	VerifyCredential(string, string) interface{}
	IsDuplicateUserName(string) bool
	IsDuplicatedEmail(string) bool
	FindUser(string) entity.User
	// EditOldPassword(models.User)
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
	db.connection.Exec(`INSERT INTO user (Email, Username, PasswordHash) VALUES (?, ?, ?)`,
		user.Email, user.Username, password)
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

func (db *userConnection) CheckIfEmailExist(username string) bool {
	var user entity.User
	db.connection.Where(&user, "username = ?", username).Take(&user)
	return user.Id != 0
}

func (db *userConnection) VerifyCredential(username string, password string) interface{} {
	var user entity.User
	res := db.connection.Where("username = ?", username).Find(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateUserName(username string) bool {
	var user entity.User
	db.connection.Where("username = ?", username).Take(&user)
	return user.Id == 0
}

func (db *userConnection) IsDuplicatedEmail(email string) bool {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user)
	return user.Id == 0
}

func (db *userConnection) FindUser(username string) entity.User {
	var user entity.User
	db.connection.Where("username = ?", username).Take(&user)
	return user
}

// func (db *userConnection) EditOldPassword(user models.User) {
// 	hash := HashPassword(user.NewPassword)
// 	db.connection.Model(entity.User{}).Where("username = ?", user.Username).
// 		UpdateColumn("PasswordHash", hash)
// }