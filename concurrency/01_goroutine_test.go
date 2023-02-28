package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func say(msg string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("say:", msg, i)
	}
}

func TestGo(t *testing.T) {
	go say("hello")
	say("world")
}
