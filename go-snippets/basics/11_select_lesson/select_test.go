package select_lesson

import (
	"errors"
	asserts "hello-world"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(10 * time.Millisecond)
		fastServer := makeDelayedServer(0)
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		expected := fastURL
		actual, err := Racer(slowURL, fastURL)

		asserts.AssertStringEquals(actual, expected, t)
		asserts.AssertNoError(err, t)
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {

		serverA := makeDelayedServer(10 * time.Millisecond)
		serverB := makeDelayedServer(11 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()

		_, actualError := ConfigurableRacer(serverA.URL, serverB.URL, 5*time.Millisecond)

		expectedError := errors.New("timeout")
		asserts.AssertErrorEquals(actualError, expectedError, t)
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
	})
}

func TestRacer2(t *testing.T) {
	slowServer := makeDelayedServer(10 * time.Millisecond)
	fastServer := makeDelayedServer(0)
	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	expected := fastURL
	actual := Racer2(slowURL, fastURL)

	asserts.AssertStringEquals(actual, expected, t)
}

func makeDelayedServer(responseDelay time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(responseDelay)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}
