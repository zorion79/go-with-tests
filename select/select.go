package _select

import (
	"errors"
	"net/http"
	"time"
)

var (
	ErrTimeout = errors.New("time is out")
	timeout    = 10 * time.Second
)

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, timeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-time.After(timeout):
		return "", ErrTimeout
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		_, _ = http.Get(url)
		close(ch)
	}()
	return ch
}
