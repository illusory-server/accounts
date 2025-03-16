package safe

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"sync/atomic"
	"testing"
)

func noop(str string) {}

func TestSafeGo(t *testing.T) {
	t.Run("Should correct work", func(t *testing.T) {
		i := int32(0)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		Go(func() {
			defer wg.Done()
			atomic.AddInt32(&i, 5)
		})
		wg.Wait()
		assert.Equal(t, int32(5), i)
	})

	t.Run("Should recovered panic", func(t *testing.T) {
		i := int32(2)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		Go(func() {
			defer wg.Done()
			var sl []string
			a := sl[0]
			noop(a)
			atomic.AddInt32(&i, 5)
		})
		wg.Wait()
		assert.Equal(t, int32(2), i)
	})
}
