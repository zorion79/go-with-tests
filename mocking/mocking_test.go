package main

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write([]byte) (int, error) {
	s.Calls = append(s.Calls, write)
	return 0, nil
}

const (
	sleep = "sleep"
	write = "write"
)

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := new(SpyTime)
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	assert.Equal(t, spyTime.durationSlept, sleepTime)
}

func TestCountdown(t *testing.T) {
	t.Run("sleep before every point", func(t *testing.T) {
		spySleepPrinter := new(CountdownOperationsSpy)

		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		assert.Equal(t, want, spySleepPrinter.Calls)
	})

	t.Run("prints 3 to GO", func(t *testing.T) {
		buffer := new(bytes.Buffer)
		Countdown(buffer, new(CountdownOperationsSpy))

		want := `3
2
1
GO!`

		assert.Equal(t, want, buffer.String())
	})
}
