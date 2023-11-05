package handlers

import (
	"encoding/json"
	"net/http"

	"come.archroid.pixelgolang/db"
	"come.archroid.pixelgolang/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	user, err := db.LoginUser(r.FormValue("email"), r.FormValue("password"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		utils.LogRequest(err.Error(), r)
	} else {
		json.NewEncoder(w).Encode(map[string]string{"token": user.AuthToken})
		utils.LogRequest("Successful", r)
	}
}
