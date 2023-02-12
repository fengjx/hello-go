package basic

import (
	"math"
	"testing"
)

// bool string
func TestBoolString(t *testing.T) {
	var (
		a = true
		b = "hello"
	)
	t.Log(a, b)
}

// int int8 int16 int32 int64
func TestInt(t *testing.T) {
	var a = 10
	var b int64
	// 显示类型转换（go 不支持隐式类型转换）
	b = int64(a)
	t.Log(a, b)
	// string 会默认赋值空字符串
	var c string
	t.Logf("%q", c)
}

func TestUint(t *testing.T) {

}

func TestMax(t *testing.T) {
	t.Log("MaxInt8:", math.MaxInt8)
	t.Log("MaxInt16:", math.MaxInt16)
	t.Log("MaxInt32:", math.MaxInt32)
	t.Log("MaxInt64:", math.MaxInt64)
	t.Log()
	t.Log("MaxUint8:", math.MaxUint8)
	t.Log("MaxUint16:", math.MaxUint16)
	t.Log("MaxUint32:", math.MaxUint32)
	t.Log()
	t.Log("MaxFloat32:", math.MaxFloat32)
	t.Log("MaxFloat64:", math.MaxFloat64)
}

// 指针
func TestPoint(t *testing.T) {
	var a = 100
	// & 为取地址符，即: a 的指针地址
	var aPoint = &a
	// *aPoint 等价于 a
	*aPoint = 300
	// go 的指针不能进行运算
	// aPoint = aPoint + 1
	t.Log(a)
}

// 自定义类型别名
type myInt int64

func TestAlias(t *testing.T) {
	var a myInt = 1024
	var b = int64(a)
	t.Log(a, b)
}
