package handlers

import (
	"db_service"
	"log"
	"net/http"
	"utils"
)

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Get statistics request was received. Starting...")
	res := db_service.GetStats()
	utils.WriteJSON(res, w, http.StatusOK)
	log.Println("Get statistics request was completed successfully")
}
