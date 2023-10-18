package basic

import (
	"fmt"
	"testing"
)

type User struct {
	name string
	age  int8
	ext  map[string]interface{}
}

// 为 User 定义方法
func (user *User) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, Ext: %v", user.name, user.age, user.ext)
}

func TestStruct(t *testing.T) {

	var u1 = &User{name: "fengjx", age: 20}
	t.Log(u1)
	t.Log(u1.age)
	t.Logf(u1.String())

	// 等价于 &User{}，返回指针类型
	var u2 = new(User)
	u2.name = "fengjx"
	u2.age = 24
	t.Log(u2)
}

type Hello interface {
	greet(name string) string
}

type HelloWorld struct {
}

func (hello *HelloWorld) greet(name string) string {
	return "hello: " + name
}

// 接口
func TestInterface(t *testing.T) {
	var hello = new(HelloWorld)
	t.Log(hello.greet("fengjx"))
}

func TestCopy(t *testing.T) {
	u1 := &User{
		name: "u1",
		age:  20,
		ext: map[string]interface{}{
			"a": "1",
		},
	}
	ext := make(map[string]interface{})
	for k, v := range u1.ext {
		ext[k] = v
	}
	u2 := u1
	func(u User) {
		u.name = "u2"
		u.age = 25
		u.ext["b"] = 2
	}(*u2)
	t.Log(u1)
}
