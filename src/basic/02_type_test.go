package basic

import (
	"fmt"
	"math"
	"testing"
	"time"
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

// 类型判断
func typeCheck(i interface{}) string {
	if v, ok := i.(int); ok {
		return fmt.Sprintf("int %d", v)
	}
	if v, ok := i.(string); ok {
		return fmt.Sprintf("string %s", v)
	}
	return "other"
}

// 类型判断，switch 写法
func typeCheck2(i interface{}) string {
	switch v := i.(type) {
	case int:
		return fmt.Sprintf("int %d", v)
	case string:
		return fmt.Sprintf("string %s", v)
	default:
		return "other"
	}
}

// 类型判断
func TestType(t *testing.T) {
	t.Log(typeCheck(100))
	t.Log(typeCheck("100"))
	t.Log(typeCheck2(200))
	t.Log(typeCheck2("200"))
}

// 日期
func TestDate(t *testing.T) {
	now := time.Now()
	t.Log(now, now.UnixMilli())
	t.Log(now.Format("2006-01-02"))
	t.Logf("%02d.%02d.%4d\n", now.Day(), now.Month(), now.Year())

	utc := time.Now().UTC()
	t.Log(utc)

	// 日期累加
	t.Log(now.Add(time.Hour * 24).Format("2006-01-02"))

	// 时间戳转日期
	dateTime := time.Unix(1676440759, 0)
	t.Log(dateTime.Format("2006-01-02 15:04:05"))
	t.Log(dateTime.UnixNano() / 1000 / 1000)
}
