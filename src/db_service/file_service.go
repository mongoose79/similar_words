package db_service

import (
	"bufio"
	"fmt"
	"log"
	"models"
	"os"
	"strings"
	"sync"
	"time"
	"utils"

	"github.com/palantir/stacktrace"
)

var InputData map[string][]string
var TotalWordsCount int
var TotalRequestsCount int
var TotalRequestsTime time.Duration

const DataInputFile = "words_clean.txt"

func InitFromFromFile() (map[string][]string, int, error) {
	log.Println(fmt.Sprintf("Start opening %s", DataInputFile))
	file, err := os.Open(DataInputFile)
	if err != nil {
		return nil, 0, stacktrace.Propagate(err, "Failed to open the source file")
	}
	log.Println(fmt.Sprintf("Opening %s was completed successfully", DataInputFile))

	totalCount := 0
	inputData := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		totalCount++
		newStr := scanner.Text()
		sortedNewStr := utils.SortString(newStr)
		if _, isStrExist := inputData[sortedNewStr]; isStrExist {
			if !utils.IsStringInSlice(newStr, inputData[sortedNewStr]) {
				inputData[sortedNewStr] = append(inputData[sortedNewStr], newStr)
			}
		} else {
			inputData[sortedNewStr] = []string{newStr}
		}
	}
	InputData = inputData
	TotalWordsCount = totalCount
	return inputData, totalCount, nil
}

func FindSimilarWords(word string, mutex *sync.Mutex) []string {
	mutex.Lock()
	defer mutex.Unlock()
	log.Println(fmt.Sprintf("Start searching similar words to '%s'", word))
	start := time.Now()
	TotalRequestsCount++
	if word == "" {
		log.Println("Input word is empty. Nothing to search")
		return nil
	}
	sortedWord := utils.SortString(word)
	var res []string
	if similarWords, isExist := InputData[sortedWord]; isExist {
		res = utils.RemoveStrFromSlice(similarWords, word)
	}
	TotalRequestsTime += time.Since(start)
	log.Println(fmt.Sprintf("Searching similar words to '%s' was completed successfully. Result: %s", word, strings.Join(res, ",")))
	return res
}

func GetStats() models.Stats {
	var avgProcTime time.Duration
	if TotalRequestsCount > 0 {
		avgProcTime = TotalRequestsTime / time.Duration(TotalRequestsCount)
	}
	log.Println(fmt.Sprintf("GetStats() was executed: TotalWords-%d, TotalRequests-%d, AvgProcessingTimeNs-%d",
		TotalWordsCount, TotalRequestsCount, avgProcTime))
	return models.Stats{TotalWords: TotalWordsCount, TotalRequests: TotalRequestsCount,
		AvgProcessingTimeNs: avgProcTime}
}

func Reset() {
	TotalRequestsCount = 0
	TotalRequestsTime = time.Duration(0)
}
