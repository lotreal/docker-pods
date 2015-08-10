package sh

import (
	"errors"
        "fmt"
	"reflect"
	"strings"
)

type Field struct {
	Start int
	End   int
	Name  string
}

type Reader struct {
	Line   int
	column int
	fields []Field
	R      []string
}

func NewReader(in []string, pattern interface{}) *Reader {
	// s  := reflect.ValueOf(pattern).Elem()
	st := reflect.TypeOf(pattern)

	column := st.NumField()
	header, in := in[0], in[1:]

	var fields []Field
	fields = make([]Field, column)

	offset := 0

	for i := 0; i < column; i++ {
		f := st.Field(i)
		colname := f.Tag.Get("field")

		index := strings.Index(header[offset:], colname)
		offset = offset + index

		fields[i] = Field{
			Start: offset,
			Name:  f.Name,
		}
	}

	for i := 0; i < (column - 1); i++ {
		fields[i].End = fields[i+1].Start
	}

	return &Reader{
		Line: len(in) - 1,
		column: column,
		fields: fields,
		R: in,
	}
}

// func TabReader(in []string) {
// 	for i := 0; i < len(in); i++ {
// 		fmt.Println(in[i])
// 	}
// 	// lines := strings.Split(string(in[:]), "\n")
// 	// fmt.Printf("in: %#v", lines)

// 	// ins := Status{}
// 	// rt  := reflect.TypeOf(ins)
// 	// fmt.Printf("%#v", rt.NumField())


// 	// fmt.Printf("in: %#v", in)
// 	// n := SliceIndex(len(in), func(i int) bool { return in[i] == 0xa })
// 	// s := string(in[:n])
// 	// fmt.Printf("line 1: %v", s)
// }

func Unmarshal(reader *Reader, record string, v interface{}) error {
	if record == "" {
		return errors.New("record is empty")
	}
	s := reflect.ValueOf(v).Elem()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)

		field := reader.fields[i]
		start := field.Start
		end   := field.End

		var value string
		if end == 0 {
			value = record[start:]
		} else {
			value = record[start:end]
		}

		f.SetString(strings.Trim(value, " "))
		fmt.Println(reader.fields[i]);
	}
	return nil
}
