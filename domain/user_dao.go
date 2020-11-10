package domain

import (
	"errors"
	"fmt"
)

var (
	users = map[int64]*UseUser{
		123: {Id: 123, FirstName: "Teste", LastName: "Outro", Email: "algo@algo.com"}
	}
)

func GetUser(userId int64) (*User, error) {

	user := users[userId]

	if user == nil {
		return nil, errors.New(fmt.Sprintf("user not found for $v", userId))
	}

	return user, nil
}
