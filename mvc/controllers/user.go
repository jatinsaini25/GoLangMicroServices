package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jatinsaini25/GoLangMicroServices/mvc/services"
	"github.com/jatinsaini25/GoLangMicroServices/mvc/utils"
)

func GetUsers(writer http.ResponseWriter, request *http.Request) {
	userId, errorInfo := strconv.ParseInt(request.URL.Query().Get("user_id"), 10, 64)

	if errorInfo != nil {
		writer.WriteHeader(http.StatusNotFound)
		errorMesg, _ := json.Marshal(&utils.ApplicationError{
			Message:   fmt.Sprintf("User Id %v should be a number", request.URL.Query().Get("user_id")),
			Code:      "Bad Request",
			ErrorCode: http.StatusBadRequest,
		})
		writer.Write(errorMesg)
		// writer.Write([]byte(error.Error))
	} else {
		userInfo, err := services.GetUsers(userId)

		if err != nil {
			detail, _ := json.Marshal(err)
			writer.Write(detail)
			return
		}
		jsonObject, _ := json.Marshal(userInfo)
		writer.Write(jsonObject)
	}

}
