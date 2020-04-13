package app

import (
	"net/http"

	"github.com/jatinsaini25/GoLangMicroServices/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUsers)

	http.ListenAndServe(":4201", nil)
}
