package handler

import "net/http"

// Handle preflight requests
func handlePreflight(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodOptions {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	} else {
		w.Write([]byte("success"))
		return
	}
}
