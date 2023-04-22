package basic

import (
	"reflect"
	"testing"
)

type UserModel struct {
}

func TestName(t *testing.T) {
	t.Log(reflect.TypeOf(UserModel{}).String())
}
