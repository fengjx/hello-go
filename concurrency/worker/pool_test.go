package worker

import (
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	p := New(5)

	for i := 0; i < 10; i++ {
		idx := i
		err := p.Schedule(func() {
			time.Sleep(time.Second * 2)
			t.Logf("task%d end", idx)
		})
		if err != nil {
			t.Log("task: ", i, "err:", err)
		}
	}
	p.Free()
}
