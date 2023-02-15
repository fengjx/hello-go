package concurrency

import (
	"testing"
	"time"
)

func listen(c1, c2, quit chan int) (int, int) {
	var s1, s2 int
	for {
		select {
		case v := <-c1:
			s1 += v
		case v := <-c2:
			s2 += v
		case <-quit:
			return s1, s2
		}
	}
}

func TestSelect(t *testing.T) {

	c1 := make(chan int)
	c2 := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c1 <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c2 <- i
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		quit <- 0
	}()
	v1, v2 := listen(c1, c2, quit)
	t.Log(v1, v2)
}
