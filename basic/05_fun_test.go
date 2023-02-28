package basic

import (
	"fmt"
	"testing"
)

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

type UserDemo struct {
	Name string
}

func (target *UserDemo) GetName() string {
	return target.Name
}

func sayHi(u *UserDemo) {
	fmt.Println("hi", u.GetName())
}

func TestUserFun(t *testing.T) {
	u := &UserDemo{
		Name: "fengjx",
	}
	sayHi(u)
}
