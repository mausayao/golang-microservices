package domain

import (
	"fmt"
	"golang-microservices/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Teste", LastName: "Outro", Email: "algo@algo.com"},
	}
)

func GetUser(userId int64) (*User, *utils.AppError) {

	user := users[userId]

	if user == nil {
		return nil, &utils.AppError{
			Message:    fmt.Sprintf("user %v was not found", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}

	return user, nil
}
