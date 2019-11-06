package syncer

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := new(Counter)
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assert.Equal(t, 3, counter.Value())
	})

	t.Run("safe concurrency", func(t *testing.T) {
		want := 1000
		counter := new(Counter)

		var wg sync.WaitGroup
		for i := 0; i < want; i++ {
			wg.Add(1)
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()
		assert.Equal(t, want, counter.Value())
	})
}
