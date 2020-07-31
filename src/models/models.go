package models

import "time"

type SimilarWords struct {
	Similar []string `json:"similar"`
}

type Stats struct {
	TotalWords          int           `json:"total_words"`
	TotalRequests       int           `json:"total_requests"`
	AvgProcessingTimeNs time.Duration `json:"avg_processing_time_ns"`
}
