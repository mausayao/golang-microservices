package controllers

import (
	"encoding/json"
	"golang-microservices/services"
	"golang-microservices/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)

	if err != nil {

		apiErr := &utils.AppError{
			Message:    "User must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bar_request",
		}

		jsonValue, _ := json.Marshal(apiErr)

		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)

	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)

		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)

	resp.Write(jsonValue)
}
