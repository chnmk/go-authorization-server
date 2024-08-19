package handler

import (
	"fmt"
	"net/http"
)

/*
This function should only be reached after the actual request handler

TODO: rewrite this function to avoid the "superfluous response.WriteHeader call" warning
in every valid non-preflight request.
*/
func handlePreflight(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unexpected request method...")

	if r.Method == http.MethodOptions {
		fmt.Println("Got preflight request!")
		w.Write([]byte("success"))
		return

	} else {
		fmt.Println("Non-preflight request in the preflight handler")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return

	}
}
