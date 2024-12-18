package _select

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, d time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(d):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(a string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(a)
		close(ch)
	}()
	return ch
}
