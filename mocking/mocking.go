package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

const (
	finalWord      = "GO!"
	countDownStart = 3
)

func main() {
	Countdown(os.Stdout)
}

func Countdown(writer io.Writer) {
	for i := countDownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		_, _ = fmt.Fprintf(writer, "%d\n", i)
	}

	time.Sleep(1 * time.Second)
	_, _ = fmt.Fprintf(writer, finalWord)
}
