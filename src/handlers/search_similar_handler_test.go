package handlers

import (
	"encoding/json"
	"fmt"
	"models"
	"net/http"
	"testing"
	"utils"

	"github.com/stretchr/testify/assert"
)

func TestSearchSimilarHandler(t *testing.T) {
	similar := invokeSimilarRequest(t, "eirsstuv")
	assert.Equal(t, []string{"revuists", "stuivers"}, similar)

	similar = invokeSimilarRequest(t, "opr")
	assert.Equal(t, []string{"por", "pro"}, similar)

	similar = invokeSimilarRequest(t, "non_existing_word")
	assert.Equal(t, []string(nil), similar)
}

func invokeSimilarRequest(t *testing.T, word string) []string {
	url := fmt.Sprintf("/similar?word=%s", word)
	request, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)

	response := utils.InvokeRequest(request, SearchSimilarHandler, "/similar")
	assert.Equal(t, http.StatusOK, response.Code)

	var similarWords models.SimilarWords
	err = json.Unmarshal(response.Body.Bytes(), &similarWords)
	assert.NoError(t, err)

	return similarWords.Similar
}
