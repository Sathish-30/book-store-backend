package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sathish-30/book-management/internal/database"
	"github.com/sathish-30/book-management/internal/model"
	"github.com/sathish-30/book-management/internal/util"
)

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"Message": "Health check",
	})
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		util.SendResponse(w, model.Response{Result: err.Error(), Status: "failure"}, http.StatusNotAcceptable)
	}
	var user model.User
	user.Name = r.FormValue("name")
	user.EmailId = r.FormValue("email_id")
	user.Password = r.FormValue("password")
	var existingUser model.User
	result := database.Db.Find(&existingUser, "email = ?", user.EmailId)
	if result.Error != nil {
		util.SendResponse(w, model.Response{Result: "Failed to retrieve data from Database", Status: "failure"}, http.StatusNotFound)
		return
	}
	if existingUser.EmailId != "" {
		util.SendResponse(w, model.Response{Result: "User already exists", Status: "failure"}, http.StatusNotFound)
		return
	}
	result = database.Db.Create(&user)
	if result.Error != nil {
		util.SendResponse(w, model.Response{Result: "Failed to add user data to the Database", Status: "failure"}, http.StatusNotFound)
		return
	}
	util.SendResponse(w, model.Response{Result: model.Content{Data: user, Message: "Added user to the data base"}, Status: "success"}, http.StatusAccepted)
}
