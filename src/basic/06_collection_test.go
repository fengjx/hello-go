package basic

import (
	"container/list"
	"reflect"
	"testing"
)

func TestArr(t *testing.T) {
	// 数组定义
	var arr1 [3]int // 每个元素会赋初始值0
	t.Log(arr1)
	var arr2 = [3]int{1, 2, 3}
	t.Log(arr2)
	var arr3 = [...]int{10, 20, 30, 40}
	t.Log(arr3)
	// 数组截取，[包含, 不包含]
	t.Log(arr3[1:3])

	// 数组访问
	t.Log(arr1[0], arr2[1], arr3[2])
	// 数组遍历
	for idx, e := range arr3 {
		t.Log(idx, e)
	}
}

func TestSlice(t *testing.T) {
	// 没有指定长度，定义的事 slice 而不是 array
	var s1 []int
	t.Log(len(s1), cap(s1))

	for i := 0; i < 20; i++ {
		s1 = append(s1, i)
		t.Log(len(s1), cap(s1))
	}
	t.Log(s1)
	var s2 = s1[13:16]
	t.Log(s2) // [13 14 15]

	var s3 = s2[1:2]
	t.Log(s3) // [14]

	var s4 = s2[1:5]
	t.Log(s4) // [14 15 16 17]，这里依然能取出数据，不会报错，因为底层的数组没有变

	var s5 = s2[1:9]
	t.Log(s5) // [14 15 16 17 18 19 0 0]，原始数组后面没有数据了，所以用 0 填充

}

// 创建 map
func TestMapInit(t *testing.T) {
	var m1 = map[string]int{"a": 1, "b": 2}
	t.Log(m1, len(m1))

	var m2 = map[string]int{}
	m2["hello"] = 1
	t.Log(m2, len(m2))

	// 初始容量
	var m3 = make(map[string]int, 5)
	t.Log(m3, len(m3))

}

func TestMapOpt(t *testing.T) {
	var m1 = map[string]int{}
	// go 会赋初始值，所以 value = 0 无法判断是值为空还是本身值就是 0
	t.Log(m1["a"])

	m1["a"] = 100
	if val, ok := m1["a"]; ok {
		t.Logf("a = %d", val)
	} else {
		t.Log("key [a] not exist")
	}
}

// 遍历 mao
func TestMapTraversal(t *testing.T) {
	var m1 = map[string]int{"a": 1, "b": 2}
	for k, v := range m1 {
		t.Logf("%s = %d", k, v)
	}
}

// 使用 Map 来实现 Set
func TestSet(t *testing.T) {
	var s = map[string]bool{}
	// 值 : 是否存在
	s["a"] = true
	s["b"] = true
	t.Log(s, len(s))

	if s["a"] {
		t.Log("[a] exist")
	} else {
		t.Log("[a] not exist")
	}

	// 删除元素
	delete(s, "b")
	t.Log(s, len(s))

	if s["b"] {
		t.Log("[b] exist")
	} else {
		t.Log("[b] not exist")
	}
}

// 链表
func TestList(t *testing.T) {

	var l = list.New()
	l.PushBack(10)
	l.PushBack("20")
	element := l.PushBack("30")

	// 从头部遍历链表
	for elm := l.Front(); elm != nil; elm = elm.Next() {
		t.Log(elm.Value, reflect.TypeOf(elm.Value))
	}

	// 删除链表元素
	l.Remove(element)
	t.Log("remove element", element.Value)

	// 从尾部遍历链表
	for elm := l.Back(); elm != nil; elm = elm.Prev() {
		t.Log(elm.Value, reflect.TypeOf(elm.Value))
	}

}
