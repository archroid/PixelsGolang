package main

import (
	"net/http"
	"os"

	"come.archroid.pixelgolang/db"
	"come.archroid.pixelgolang/handlers"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/color"
)

func main() {

	//create userdata folder
	os.MkdirAll("userdata", os.ModePerm)

	db.Init()

	router := mux.NewRouter()

	router.HandleFunc("/ping", handlers.PingHandler).Methods("GET")
	router.HandleFunc("/user/create", handlers.CreateUserHandler).Methods("POST")
	router.HandleFunc("/user/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/post/create", handlers.CreatePostHandler).Methods("POST")
	router.HandleFunc("/post/getbyID", handlers.GetPostByIdHandler).Methods("POST")

	staticDir := "/userdata/"
	http.Handle(staticDir, http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	http.Handle("/", router)
	color.Println(color.Blue("listening on port 5000"))

	http.ListenAndServe(":5000", nil)
}
