package handlers

import (
	"encoding/json"
	"net/http"

	"come.archroid.pixelgolang/db"
	"come.archroid.pixelgolang/utils"
)

func GetPostByIdHandler(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")

	post, err := db.GetPostById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		utils.LogRequest(err.Error(), r)
		return
	}

	json.NewEncoder(w).Encode(post)
	utils.LogRequest("Successful", r)

}
