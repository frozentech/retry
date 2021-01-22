package retry

import (
	"math/rand"
	"time"
)

// Retrier defines a Retryer function
type Retrier struct {
	// Loop defines the number of retries if an error occurs
	Loop int

	// Delay defines a delay function that returns time.Duration
	Delay func() time.Duration
}

// NoDelay ...
func NoDelay() time.Duration { return 0 }

// RandomDelay ...
func RandomDelay() time.Duration { return time.Duration(rand.Intn(100)) * time.Millisecond }

// Try executes `fn` and retries on error
func (me Retrier) Try(fn func() error) (err error) {
	var (
		n int
	)

	for n < me.Loop {
		err := fn()
		if err == nil {
			return nil
		}

		n++
		time.Sleep(me.Delay())
	}

	return
}
