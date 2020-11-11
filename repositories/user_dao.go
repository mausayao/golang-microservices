package repositories

import (
	"fmt"
	"golang-microservices/domain"
	"golang-microservices/utils"
	"net/http"
)

var (
	users = map[int64]*domain.User{
		123: {Id: 123, FirstName: "Teste", LastName: "Outro", Email: "algo@algo.com"},
	}
)

type UserDao struct {
}

func (c UserDao) GetUser(userId int64) (domain.User, *utils.AppError) {

	user := users[userId]

	if user == nil {
		return domain.User{}, &utils.AppError{
			Message:    fmt.Sprintf("user %v was not found", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}

	return *user, nil

}
