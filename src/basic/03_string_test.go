package basic

import (
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestStr(t *testing.T) {
	var str = "hello"
	t.Log(str[0], str[3])

	var a = `
import (
	"testing"
	"unicode/utf8"
)
`
	t.Log(a)
}

// 字符长度
func TestLen(t *testing.T) {
	var str = "你好"
	t.Log(len(str))
	t.Log(utf8.RuneCountInString(str))

	var a = "hello"
	t.Log(len(a))
	t.Log(utf8.RuneCountInString(a))

	// 表情占4个字符
	var b = "😄"
	t.Log(len(b))
	t.Log(utf8.RuneCountInString(b))
}

func TestFormat(t *testing.T) {
	var a = "你好"
	var r = []rune(a)
	t.Logf("Unicode - %s: %x", a, r[0])
	t.Logf("UTF8 - %[1]s: %[1]x", a)
	for _, c := range a {
		t.Logf("%[1]c %[1]x", c)
	}
}

func TestUtil(t *testing.T) {
	var str = "a,b,c"
	var arr = strings.Split(str, ",")
	t.Log("1=>", arr)
	t.Log("2=>", strings.Join(arr, "-"))

	var i = 1024
	// int to string
	var is = strconv.Itoa(i)
	t.Log("3=>", is)

	var s = "1024"
	var si, _ = strconv.Atoi(s)
	t.Log("4=>", si+1000)
}
