package sh

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"reflect"
)

func InterfaceSlice(in interface{}) []interface{} {
	s := reflect.ValueOf(in)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i:=0; i<s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func Struct2Array(x interface{}) []string {
	v := reflect.ValueOf(x)

	values := make([]string, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = string(v.Field(i).Interface().(string))
	}

	return values
}

func StructKeys(x interface{}) []string {
	v := reflect.TypeOf(x)
	keys := make([]string, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		colname := v.Field(i).Tag.Get("field")
		keys[i] = colname
	}

	return keys
}

func TabWrite(in interface{}) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 5, 8, 1, '\t', tabwriter.AlignRight)

	v := InterfaceSlice(in)
	header := strings.Join(StructKeys(v[0]), "\t")
	fmt.Fprintln(w, header)

	for _, val := range v {
		a := Struct2Array(val)
		fmt.Fprintln(w, strings.Join(a, "\t"))
	}

	fmt.Fprintln(w)
	w.Flush()
}
