package context1

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store got canceled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func TestHandler(t *testing.T) {
	data := "Hello, world!"
	t.Run("return data from store", func(t *testing.T) {
		stubStore := new(SpyStore)
		stubStore.response = data
		stubStore.t = t
		srv := Server(stubStore)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		assert.Equal(t, data, response.Body.String())
	})

	t.Run("test with cancel", func(t *testing.T) {
		store := new(SpyStore)
		store.response = data
		store.t = t
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancelCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancelCtx)

		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}
