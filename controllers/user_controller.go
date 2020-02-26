package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rennanbadaro/go-mvc-webserver/services"
	"github.com/rennanbadaro/go-mvc-webserver/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIDParam := req.URL.Query().Get("id")
	userID, err := strconv.ParseInt(userIDParam, 10, 64)

	if err != nil {
		apiErr := utils.ApplicationError{
			StatusCode: http.StatusBadRequest,
			Message:    "ID must be a number",
			Code:       "bad_req",
		}

		parsedErr, _ := json.Marshal(apiErr)

		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(parsedErr)

		return
	}

	user, apiErr := services.GetUser(userID)

	if apiErr != nil {
		parsedErr, _ := json.Marshal(apiErr)

		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(parsedErr)

		return
	}

	parsedUser, _ := json.Marshal(user)

	resp.WriteHeader(http.StatusOK)
	resp.Write(parsedUser)
}
