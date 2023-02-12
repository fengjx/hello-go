package basic

import "testing"

func hello(name string) string {
	return "hello: " + name
}

func TestHello(t *testing.T) {
	t.Log(hello("fengjx"))
}

// 可变参数
func sum(nums ...int) int {
	var s = 0
	for _, n := range nums {
		s += n
	}
	return s
}

func TestSum(t *testing.T) {
	t.Log(sum(1, 2, 10))
}
