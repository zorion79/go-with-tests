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
	canceled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.canceled = true
}

func TestHandler(t *testing.T) {
	data := "Hello, world!"
	t.Run("return data from store", func(t *testing.T) {
		stubStore := new(SpyStore)
		stubStore.response = data
		srv := Server(stubStore)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		assert.Equal(t, data, response.Body.String())
		assert.False(t, stubStore.canceled, "should be canceled")
	})

	t.Run("test with cancel", func(t *testing.T) {
		store := new(SpyStore)
		store.response = data
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancelCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancelCtx)

		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		assert.True(t, store.canceled, "store not canceled")
	})
}
