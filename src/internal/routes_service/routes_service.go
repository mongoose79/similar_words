package routes_service

import (
	"internal/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/palantir/stacktrace"
)

func InitRoutes() error {
	log.Println("Configuring routes")
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1/").Subrouter()
	subRouter.HandleFunc("/similar", handlers.SearchSimilarHandler)
	subRouter.HandleFunc("/stats", handlers.StatsHandler)
	http.Handle("/", router)

	log.Println("Server is listening in the port 8000...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return stacktrace.Propagate(err, "Failed to init the routes")
	}
	return nil
}
