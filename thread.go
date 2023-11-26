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

type JoinHandle struct {
	channel chan Unit
}

func (self JoinHandle) Join() Result[Unit] {
	<-self.channel

	return Ok(Unit{})
}

func Spawn(f func()) JoinHandle {
	channel := make(chan Unit)

	go func() {
		f()
		channel <- Unit{}
	}()

	return JoinHandle{
		channel,
	}
}
