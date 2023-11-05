package handlers

import (
	"encoding/json"
	"net/http"

	"come.archroid.pixelgolang/utils"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(map[string]string{"status": "pong"})

	utils.LogRequest("Pong", r)
}
