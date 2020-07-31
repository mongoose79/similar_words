package handlers

import (
	"encoding/json"
	"fmt"
	"internal/db_service"
	"internal/models"
	"internal/utils"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func init() {
	db_service.InitFromFromFile()
}

func TestStatsHandlerNoRequests(t *testing.T) {
	db_service.Reset()

	request, err := http.NewRequest("GET", "/stats", nil)
	assert.NoError(t, err)
	response := utils.InvokeRequest(request, StatsHandler, "/stats")
	assert.Equal(t, http.StatusOK, response.Code)

	var stats models.Stats
	err = json.Unmarshal(response.Body.Bytes(), &stats)
	assert.NoError(t, err)

	assert.Equal(t, 351075, stats.TotalWords)
	assert.Equal(t, 0, stats.TotalRequests)
	assert.Equal(t, time.Duration(0), stats.AvgProcessingTimeNs)
}

func TestStatsHandler(t *testing.T) {
	db_service.Reset()
	_, _, err := db_service.InitFromFromFile()
	assert.NoError(t, err)

	similar := invokeSimilarRequest(t, "eirsstuv")
	assert.Equal(t, []string{"revuists", "stuivers"}, similar)

	similar = invokeSimilarRequest(t, "opr")
	assert.Equal(t, []string{"por", "pro"}, similar)

	similar = invokeSimilarRequest(t, "non_existing_word")
	assert.Equal(t, []string(nil), similar)

	request, err := http.NewRequest("GET", "/stats", nil)
	assert.NoError(t, err)
	response := utils.InvokeRequest(request, StatsHandler, "/stats")
	assert.Equal(t, http.StatusOK, response.Code)

	var stats models.Stats
	err = json.Unmarshal(response.Body.Bytes(), &stats)
	assert.NoError(t, err)

	assert.Equal(t, 351075, stats.TotalWords)
	assert.Equal(t, 3, stats.TotalRequests)
	assert.Equal(t, time.Duration(0), stats.AvgProcessingTimeNs)
}

func TestStatsHandlerParallel(t *testing.T) {
	db_service.Reset()
	start := time.Now()
	iter := 100000
	var wg sync.WaitGroup
	wg.Add(iter)
	for i := 0; i < iter; i++ {
		go invokeSimilarRequestParallel(t, "eirsstuv", &wg)
	}
	wg.Wait()
	duration := time.Since(start)
	str := fmt.Sprintf("%d", int(duration.Seconds()))
	fmt.Println(str)

	request, err := http.NewRequest("GET", "/stats", nil)
	assert.NoError(t, err)
	response := utils.InvokeRequest(request, StatsHandler, "/stats")
	assert.Equal(t, http.StatusOK, response.Code)

	var stats models.Stats
	err = json.Unmarshal(response.Body.Bytes(), &stats)
	assert.NoError(t, err)

	assert.Equal(t, 351075, stats.TotalWords)
	assert.Equal(t, iter, stats.TotalRequests)
}

func invokeSimilarRequestParallel(t *testing.T, word string, wg *sync.WaitGroup) {
	url := fmt.Sprintf("/similar?word=%s", word)
	request, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)

	response := utils.InvokeRequest(request, SearchSimilarHandler, "/similar")
	assert.Equal(t, http.StatusOK, response.Code)

	wg.Done()
}
