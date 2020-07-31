package handlers

import (
	"internal/db_service"
	"internal/models"
	"internal/utils"
	"log"
	"net/http"
	"sync"
)

func SearchSimilarHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Search similar words request was received. Starting...")
	word := r.URL.Query().Get("word")
	if word == "" {
		errMsg := "Failed source argument 'word'"
		log.Println(errMsg)
		utils.WriteJSON(errMsg, w, http.StatusBadRequest)
		return
	}
	var res models.SimilarWords
	var mutex sync.Mutex
	res.Similar = db_service.FindSimilarWords(word, &mutex)
	utils.WriteJSON(res, w, http.StatusOK)
	log.Println("Search similar words request was completed successfully")
}
