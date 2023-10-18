package basic

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type UserModel struct {
}

func TestName(t *testing.T) {
	t.Log(reflect.TypeOf(UserModel{}).String())
}

var types = map[string]reflect.Type{}

func Register(value interface{}) string {
	t := reflect.TypeOf(value).Elem()
	name := fmt.Sprintf("%s.%s", t.PkgPath(), t.Name())
	types[name] = t
	return name
}

func TypeByName(name string) reflect.Type {
	return types[name]
}

type M1 struct {
	Uid      int64  `json:"uid"`
	Nickname string `json:"nickname"`
}

type DataWrap struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func TestJson(t *testing.T) {
	m1 := &M1{
		Uid:      100,
		Nickname: "hello",
	}
	m1JsonByt, _ := json.Marshal(m1)
	dw := &DataWrap{
		Type: Register(m1),
		Data: string(m1JsonByt),
	}
	jsonByt, _ := json.Marshal(dw)
	jsonStr := string(jsonByt)
	t.Log(jsonStr)

	dw2 := &DataWrap{}
	_ = json.Unmarshal(jsonByt, dw2)
	vp := reflect.New(TypeByName(dw2.Type))
	// v := reflect.Indirect(vp)
	err := json.Unmarshal([]byte(dw2.Data), vp.Interface())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(vp.Interface().(*M1).Uid)
}
