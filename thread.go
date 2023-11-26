package gost

import (
	"fmt"
	"time"
)

// Puts the current thread to sleep for at least the specified amount of time.
//
//	gost.Sleep(gost.DurationFromSecs(5)) // Sleep for 5 seconds
func Sleep(dur Duration) {
	durationString := fmt.Sprintf("%ds%dns", dur.seconds, dur.nanoseconds)

	duration, err := time.ParseDuration(durationString)

	if err != nil {
		panic(err)
	}

	time.Sleep(duration)
}
