package routers

import (
    "net/http"
    "github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()

    // Define your routes here
    r.HandleFunc("/", HomeHandler).Methods("GET")

    return r
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Welcome to The Sphere Online"))
}
