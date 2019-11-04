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

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const (
	finalWord      = "GO!"
	countDownStart = 3
)

func main() {
	sleeper := new(ConfigurableSleeper)
	sleeper.duration = 1 * time.Second
	sleeper.sleep = time.Sleep

	Countdown(os.Stdout, sleeper)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		sleeper.Sleep()
		_, _ = fmt.Fprintln(writer, i)
	}

	sleeper.Sleep()
	_, _ = fmt.Fprint(writer, finalWord)
}
