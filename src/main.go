package main

import (
	"db_service"
	"fmt"
	"io"
	"log"
	"os"
	"routes_service"
)

const logFile = "SimilarWords_task.log"

func main() {
	initLog()
	log.Println("Starting Similar Words task...")
	_, totalWords, err := db_service.InitFromFromFile()
	if err != nil {
		log.Println(fmt.Sprintf("Failed to init the data from file: %v", err))
	}
	log.Println(fmt.Sprintf("Input data was loaded successfully. Total files are %d", totalWords))

	err = routes_service.InitRoutes()
	if err != nil {
		log.Println(fmt.Sprintf("Failed to init the routes: %v", err))
	}
}

func initLog() {
	fmt.Println("Start initializing the log")
	logFile, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Failed to create log file")
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}
