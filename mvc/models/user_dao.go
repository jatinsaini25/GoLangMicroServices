package models

import (
	"fmt"
	"net/http"

	"github.com/jatinsaini25/GoLangMicroServices/mvc/utils"
)

var (
	user = map[int64]*User{
		123: {Id: 123, FirstName: "Bob", LastName: "Thruston"},
	}
)

func GetUser(user_id int64) (*User, *utils.ApplicationError) {
	fmt.Println("User details : ", user[user_id])
	if userInfo := user[user_id]; userInfo == nil {
		fmt.Println("UserInfo : ", userInfo)
		errorData := &utils.ApplicationError{
			Message:   fmt.Sprintf("User %d was not found", user_id),
			ErrorCode: http.StatusNotFound,
			Code:      "not_found",
		}
		return nil, errorData
	}

	return user[user_id], nil
}
