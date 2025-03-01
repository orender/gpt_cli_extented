package utils

import "log"

// LogError logs errors with a consistent format
func LogError(err error) {
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}
