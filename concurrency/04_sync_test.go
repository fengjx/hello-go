package concurrency

import (
	"sync"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	var mtx sync.Mutex

	sum := 0
	for i := 0; i < 1000; i++ {
		go func() {
			defer mtx.Unlock()
			mtx.Lock()
			sum++
		}()
	}
	time.Sleep(3 * time.Second)
	t.Log(sum)
}
