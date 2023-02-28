package basic

import (
	"fmt"
	"os"
	"runtime"
	"testing"
	"unicode/utf8"
)

func TestCondition(t *testing.T) {
	var fileByte, err = os.ReadFile("./01_var_test.go")
	if err != nil {
		t.Log("err", err)
		return
	}
	var str = string(fileByte[:])
	if utf8.RuneCountInString(str) > 100 {
		t.Log("file content to long")
	} else {
		t.Log("press")
	}
}

func TestSwitch(t *testing.T) {
	// 在 golang 中，case 不需要加 break
	switch platform := runtime.GOOS; platform {
	case "darwin":
		t.Log("macOS.")
	case "linux":
		t.Log("Linux.")
	default:
		t.Log("not support")
	}

}

func TestLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}

func TestLoopRange(t *testing.T) {

	str := "golang"
	for idx, s := range str {
		t.Log(idx, string(s))
	}

}

func TestFori(t *testing.T) {
	for i := 1; i < 99; i++ {
		var str string
		for j := 1; j < 99; j++ {
			if j > i {
				break
			}
			str += fmt.Sprintf("%d * %d = %d, ", j, i, i*j)
		}
		t.Log(str)
	}
}

// 死循环
func TestForever(t *testing.T) {
	n := 0
	for {
		n++
		if n > 200 {
			break
		}
	}
	t.Log("end")
}

func TestDefer(t *testing.T) {

	defer t.Logf("finaly")

	t.Log("do something")
}

func TestDeferPanic(t *testing.T) {

	defer t.Logf("finally")

	t.Log("do something")
	// 出现异常也会执行 defer，类似 java 的 try catch finally
	panic("error")
}

func TestDeferStack(t *testing.T) {
	// 推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用
	for i := 0; i < 5; i++ {
		defer t.Log("defer", i)
	}
	t.Log("do something")
}
