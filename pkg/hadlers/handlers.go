package hadlers

import (
	"fmt"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Bameg")

}
