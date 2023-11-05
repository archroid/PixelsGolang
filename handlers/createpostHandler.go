package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"come.archroid.pixelgolang/db"
	"come.archroid.pixelgolang/utils"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {

	id := utils.GenerateSecureToken(40)

	// download and save image
	file, _, err := r.FormFile("file")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	os.MkdirAll("userdata/postimages", os.ModePerm)
	f, err := os.OpenFile("userdata/postimages/"+id, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, _ = io.Copy(f, file)

	imageurl := "http://192.168.1.100:5000/userdata/postimages/" + id

	post, err := db.CreatePost(r.FormValue("token"), imageurl, r.FormValue("caption"), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		utils.LogRequest(err.Error(), r)
	} else {
		json.NewEncoder(w).Encode(map[string]string{"id": post.ID})
		utils.LogRequest("Successful", r)
	}

}
