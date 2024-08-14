package routers

import (
	"github.com/gorilla/mux"
	"github.com/thesphereonline/chain/backend/api/handlers"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/user", handlers.GetUser).Methods("GET")
	return router
}