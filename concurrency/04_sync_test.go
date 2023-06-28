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

func TestTaskWait(t *testing.T) {
	ch := make(chan struct{})
	task := func() {
		time.Sleep(time.Second * 1)
		ch <- struct{}{}
	}
	go func() {
		task()
	}()
	select {
	case <-ch:
		t.Log("ok")
	case <-time.After(time.Second * 2):
		t.Log("time out")
	}
}
