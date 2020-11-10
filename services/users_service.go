package services

import (
	"golang-microservices/domain"
	"golang-microservices/utils"
)

func GetUser(userId int64) (*domain.User, *utils.AppError) {
	return domain.GetUser(userId)
}
