package concurrency

import (
	"errors"
	"fmt"
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

func TestSelectChannelBuff(t *testing.T) {
	ch := make(chan int64, 1)
	quit := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			ch <- int64(i)
			t.Logf("send %d", i)
		}
	}()

	go func() {
		for {
			select {
			case v := <-ch:
				t.Logf("rec %d", v)
				time.Sleep(time.Second * 1)
				if v == 9 {
					quit <- struct{}{}
				}
			}
		}
	}()
	<-quit
	t.Log("exit")
}

// 使用 channel 实现对象池
type ObjectPool struct {
	ch chan interface{}
}

func (target *ObjectPool) Get(timeout time.Duration) (interface{}, error) {
	select {
	case obj := <-target.ch:
		return obj, nil
	case <-time.After(timeout):
		return nil, errors.New("get obj timeout")
	}
}

func (target *ObjectPool) Release(obj interface{}) error {
	select {
	case target.ch <- obj:
		fmt.Printf("release obj: %x \n", obj)
	default:
		return errors.New("release obj error")
	}
	return nil
}
