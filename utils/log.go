package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/gommon/color"
)

func LogRequest(message string, r *http.Request) {

	dt := time.Now().Format("01-02-2006 15:04:05")

	color.Print(color.Red(r.RequestURI))
	fmt.Print(" ")

	color.Print(color.Blue(r.Method))
	fmt.Print(" ")

	color.Print(color.White(dt))

	color.Println(color.Green(" ==> " + message))
}
