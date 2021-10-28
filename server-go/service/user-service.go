package service

import (
	"github.com/adinovcina/entity"
	"github.com/adinovcina/models"
	"github.com/adinovcina/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(entity.User) entity.User
	IsDuplicatedEmail(string) bool
	VerifyCredential(string, string) interface{}
	IsOldPasswordCorrect(models.UserProfile) bool
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRep repository.UserRepository) UserService {
	return &userService{
		userRepository: userRep,
	}
}

func (service *userService) CreateUser(user entity.User) entity.User {
	res := service.userRepository.InsertUser(user)
	return res
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (service *userService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparedPassword := CheckPasswordHash(password, v.Passwordhash)
		if v.Email == email && comparedPassword {
			return entity.User{Id: v.Id, Email: v.Email, FirstName: v.FirstName, LastName: v.LastName}
		}
	}
	return false
}

func (service *userService) IsDuplicatedEmail(email string) bool {
	res := service.userRepository.IsDuplicatedEmail(email)
	return res
}

func (service *userService) IsOldPasswordCorrect(user models.UserProfile) bool {
	res := service.userRepository.FindUser(user.Email)
	isPasswordCorrect := CheckPasswordHash(user.OldPassword, res.Passwordhash)
	if isPasswordCorrect {
		service.userRepository.EditOldPassword(user)
		return true
	}
	return false
}
