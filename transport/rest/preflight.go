package handler

import "net/http"

/*
	TODO: rewrite this function
	to avoid "superfluous response.WriteHeader call" warning
	in every valid non-preflight request.
*/
func handlePreflight(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodOptions {
		http.Error(w, "Method not allowed lmao", http.StatusMethodNotAllowed)
		return
	} else {
		w.Write([]byte("success"))
		return
	}
}
