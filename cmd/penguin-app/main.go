package main

import (
	"net/http"

	"github.com:gabxway/penguin-api.git/pkg/handlers"
)

func main() {
	runAPI()
}

func runAPI() {
	http.HandleFunc("/test", handlers.TestHandler)
	http.ListenAndServe(":8081", nil)

}
