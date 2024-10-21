package select_lesson

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timeout")
	}
}

// chan struct{} is the smallest data type available from a memory perspective so we get no allocation versus a bool or any other primitive
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func Racer2(a, b string) (winner string) {
	aDuration := getResponseTime(a)
	bDuration := getResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func getResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	responseTime := time.Since(start)
	return responseTime
}
