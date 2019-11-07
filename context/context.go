package context1

import (
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

		data := make(chan string)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case <-ctx.Done():
			store.Cancel()
		case d := <-data:
			_, _ = fmt.Fprintf(w, d)
		}
	}
}
