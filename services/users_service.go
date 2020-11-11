package services

import (
	"golang-microservices/domain"
	"golang-microservices/repositories"
	"golang-microservices/utils"
)

type UserService struct{}

func (s UserService) GetUser(userId int64) (*domain.User, *utils.AppError) {
	repository := repositories.UserDao{}
	return repository.GetUser(userId)
}
