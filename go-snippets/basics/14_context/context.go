package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}

// Second implementation, for which "We can see after this that the server code has become
// simplified as it's no longer explicitly responsible for cancellation, it simply passes
// through context and relies on the downstream functions to respect any cancellations that
// may occur."

type Store2 interface {
	Fetch(ctx context.Context) (string, error)
}

func Server2(store Store2) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			fmt.Println(err.Error())
			return // todo: log error however you like
		}

		fmt.Fprint(w, data)
	}
}
