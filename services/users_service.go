package services

import "golang-microservices/domain"

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
