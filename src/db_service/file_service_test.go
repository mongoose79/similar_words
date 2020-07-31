package db_service

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInitFromFromFile(t *testing.T) {
	_, totalWords, err := InitFromFromFile()
	assert.NoError(t, err)
	assert.Equal(t, 351075, totalWords)
}

func TestGetStats(t *testing.T) {
	testing.Init()
	_, _, err := InitFromFromFile()
	assert.NoError(t, err)

	var mutex sync.Mutex
	res := FindSimilarWords("eirsstuv", &mutex)
	assert.Equal(t, []string{"revuists", "stuivers"}, res)

	res = FindSimilarWords("", &mutex)
	assert.Nil(t, res)

	res = FindSimilarWords("iersstuv", &mutex)
	assert.Equal(t, []string{"revuists", "stuivers"}, res)

	resStats := GetStats()
	assert.Equal(t, 351075, resStats.TotalWords)
	assert.Equal(t, 3, resStats.TotalRequests)
	assert.Equal(t, time.Duration(0), resStats.AvgProcessingTimeNs)
}

func TestFindSimilarWords(t *testing.T) {
	_, _, err := InitFromFromFile()
	assert.NoError(t, err)

	var mutex sync.Mutex
	res := FindSimilarWords("eirsstuv", &mutex)
	assert.Equal(t, []string{"revuists", "stuivers"}, res)
}

func TestFindSimilarWordsEmptyWord(t *testing.T) {
	var mutex sync.Mutex
	res := FindSimilarWords("", &mutex)
	assert.Nil(t, res)
}
