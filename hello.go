package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := strings.Split("http://hising.com/user/info?a=1", "?")
	fmt.Println(arr)
}
