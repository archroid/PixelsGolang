package handlers

import (
	"encoding/json"
	"net/http"

	"come.archroid.pixelgolang/db"
	"come.archroid.pixelgolang/utils"
)

func GetUserPostsHandler(w http.ResponseWriter, r *http.Request) {

	posts, err := db.GetUserPosts(r.FormValue("token"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		utils.LogRequest(err.Error(), r)
		return
	}

	json.NewEncoder(w).Encode(posts)
	utils.LogRequest("Successful", r)

}
