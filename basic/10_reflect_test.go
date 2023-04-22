package basic

import (
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	sli := []string{"a", "b", "c", "d"}
	value := reflect.ValueOf(&sli)
	t.Log(value.Type().Elem().Elem().Kind())
	t.Log(reflect.TypeOf(sli).Elem().Kind())
}

type User2 struct {
	Id string
}

func TestReflect2(t *testing.T) {
	sli := []User2{
		User2{Id: "1000"},
	}
	value := reflect.ValueOf(&sli)
	t.Log(value.Type())
	t.Log(value.Type().Elem().Kind())
	t.Log(value.Type().Elem().Elem())
	t.Log(value.Elem().Kind())
}
