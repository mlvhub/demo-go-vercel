package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mlvhub/demo-go-vercel/api/users"
)

func API(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", users.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/users", users.FetchAll).Methods(http.MethodGet)
	router.HandleFunc("/api/users/{id}", users.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/users/{id}", users.FetchByID).Methods(http.MethodGet)
	router.HandleFunc("/api/users/{id}", users.Remove).Methods(http.MethodDelete)

	router.ServeHTTP(w, r)
}
