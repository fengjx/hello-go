package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func sum(arr []int, ch chan int) {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	ch <- sum
}

func TestChannel(t *testing.T) {
	ch := make(chan int)
	go sum([]int{10, 20, 30}, ch)
	go sum([]int{100, 200, 300}, ch)
	x, y := <-ch, <-ch
	t.Log(x, y, x+y)
}

// channel 缓存
func TestChannelBuffer(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	go func() {
		t.Log("chan 3 start")
		// 缓存区满后会阻塞
		ch <- 3
		t.Log("chan 3 end")
	}()
	time.Sleep(3 * time.Second)
	t.Log("sleep end")
	t.Log(<-ch)
	t.Log(<-ch)
	t.Log(<-ch)
}

// https://tour.go-zh.org/concurrency/4
func TestChannelClose(t *testing.T) {
	ch := make(chan int)
	go func() {
		ch <- 1024
		close(ch)
	}()
	time.Sleep(1 * time.Second)
	//t.Log(<-ch)
	// 调用 close 后也会等前面的数据取出后才会返回 false，可以去掉上面的注释来测试下
	if v, ok := <-ch; ok {
		t.Log(v)
	} else {
		t.Log("channel closed")
	}
}

// 循环 for v := range ch 会不断从信道接收值，直到它被关闭
func TestChannelCloseRange(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(500 * time.Millisecond)
		}
		close(ch)
	}()
	for v := range ch {
		t.Log(v)
	}
}

// 实现生产者和消费者

type Producer struct {
	ch chan interface{}
}

func (target Producer) start() {
	for i := 0; i < 10; i++ {
		target.ch <- i
	}
}

type Consumer struct {
	ch chan interface{}
}

func (target Consumer) start() {
	for msg := range target.ch {
		fmt.Println(msg)
	}
}

func TestMsg(t *testing.T) {
	ch := make(chan interface{})
	p := &Producer{
		ch: ch,
	}
	go p.start()

	c := &Consumer{
		ch: ch,
	}
	go c.start()
	time.Sleep(3 * time.Second)
}
