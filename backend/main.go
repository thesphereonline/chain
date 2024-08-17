package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thesphereonline/chain/backend/routers"
)

func main() {
	router := mux.NewRouter()
	router = routers.SetUserRouters(router)

	log.Println("Server listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
