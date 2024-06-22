package basic

import (
	"os"
	"reflect"
	"testing"
	"text/template"
)

func TestTemplate(t *testing.T) {

	var fns = template.FuncMap{
		"last": func(x int, a interface{}) bool {
			return x == reflect.ValueOf(a).Len()-1
		},
	}

	tmp := template.Must(template.New("abc").Funcs(fns).Parse(`{{range  $i, $e := .items}}{{if $i}}, {{end}}{{if last $i $}}and {{end}}{{$e.Name}}{{end}}.`))

	type item struct {
		Name string
		Val  int64
	}

	data := map[string][]item{
		"items": {
			{
				Name: "a1",
				Val:  1,
			},
			{
				Name: "a2",
				Val:  2,
			},
		},
	}
	err := tmp.Execute(os.Stdout, data)
	if err != nil {
		t.Error(err)
	}
}
