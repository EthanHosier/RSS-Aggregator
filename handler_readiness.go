package main

import "net/http"

// response with 200 if server alive and running
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
