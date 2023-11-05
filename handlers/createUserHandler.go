package handlers

import (
	"encoding/json"
	"net/http"

	"come.archroid.pixelgolang/db"
	"come.archroid.pixelgolang/utils"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	user, err := db.CreateUser(r.FormValue("username"), r.FormValue("password"), r.FormValue("email"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		utils.LogRequest(err.Error(), r)
	} else {
		json.NewEncoder(w).Encode(map[string]string{"token": user.AuthToken})
		utils.LogRequest("Successful", r)
	}

}
