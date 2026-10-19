package selectt

import (
	"errors"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", errors.New("more than 10 seconds")
	}
}

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
		}
		close(ch)
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	resp, err := http.Get(url)
	if err == nil {
		resp.Body.Close()
	}
	return time.Since(start)
}
