package basic

import (
	"testing"
)

// 变量定义
func TestVar(t *testing.T) {
	// 类型推断
	var a = 1
	// 指定变量类型
	var b int = 1
	var aa, bb = 10, 20

	var (
		c = "hello"
		d = 100
	)

	// 省略 var
	e := "world"

	t.Log(a)
	t.Log(b)
	t.Logf("aa=%d, bb=%d", aa, bb)
	t.Log(c)
	t.Log(d)
	t.Log(e)
}

// 交换两个变量的值
func TestSwap(t *testing.T) {
	var (
		a = 10
		b = 20
	)
	t.Logf("a=%d, b=%d", a, b)
	// 交换两个变量的值
	a, b = b, a
	t.Logf("a=%d, b=%d", a, b)
}

// 常量
func TestConst(t *testing.T) {
	const (
		Monday = 1 + iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
	t.Log(Monday, Tuesday, Wednesday, Tuesday, Friday, Saturday, Sunday)

	const (
		No = iota
		Yes
	)
	t.Log(No, Yes)

	const (
		Readable = 1 << iota
		Writeable
		Executable
	)
	t.Log(Readable, Writeable, Executable)
	var flag = 7
	t.Log(flag&Readable == Readable, flag&Writeable == Writeable, flag&Executable == Executable)

	const who = "fengjx"
	t.Log(who)
}
