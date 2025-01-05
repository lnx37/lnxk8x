package util

import (
	"log"
	"time"
)

func TimeTaken(started_at time.Time) {
	var time_elapsed time.Duration
	time_elapsed = time.Since(started_at)
	log.Printf("time elapsed: %s\n", time_elapsed)
}
