package services

import (
	"github.com/jatinsaini25/GoLangMicroServices/mvc/models"
	"github.com/jatinsaini25/GoLangMicroServices/mvc/utils"
)

func GetUsers(userId int64) (*models.User, *utils.ApplicationError) {
	return models.GetUser(userId)
}