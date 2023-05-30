package main

import (
	"net/http"
	"github.com/ciscoFran10/energy-api/handlers"
)


func main() {
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/consumption", getConsumption)
	http.ListenAndServe(":8080", addCorsHeaders(http.DefaultServeMux))
}


func addCorsHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		next.ServeHTTP(w, r)
	})
}
