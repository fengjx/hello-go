package basic

import (
	"fmt"
	"testing"
)

type User struct {
	name string
	age  int8
}

// 为 User 定义方法
func (user *User) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", user.name, user.age)
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
