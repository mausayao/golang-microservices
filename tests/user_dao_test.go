package tests

import (
	"golang-microservices/domain"
	"golang-microservices/repositories"
	"golang-microservices/utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func loadData(id int64) (*domain.User, *utils.AppError) {
	repository := repositories.UserDao{}
	user, err := repository.GetUser(id)
	return user, err
}

func TestGetUserNoUserFound(t *testing.T) {
	
	user, err := loadData(0)

	assert.Nil(t, user, "we werer not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, err.StatusCode, http.StatusNotFound)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t,"user 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := loadData(123)

	assert.Nil(t, err)
	assert.NotNil(t,user)

	assert.Equal(t, int64(123), user.Id)

	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Teste", user.FirstName)
	assert.EqualValues(t, "Outro", user.LastName)
	assert.EqualValues(t, "algo@algo.com", user.Email)
}
