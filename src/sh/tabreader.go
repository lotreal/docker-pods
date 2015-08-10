package sh

import (
	"errors"
	"io"
	"reflect"
	"strings"
)

type Field struct {
	Start int
	End   int
	Name  string
}

type Reader struct {
	pointer int
	Line    int
	column  int
	fields  []Field
	R       []string
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

func (r *Reader) Read() (record string, err error) {
	record = ""
	if r.pointer <= r.Line {
		record = r.R[r.pointer]
		r.pointer++
		return record, nil
	}
	return record, io.EOF
}

func Unmarshal(reader *Reader, v interface{}) error {
	record, err := reader.Read()
	if err != nil {
		return err
	}
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
	}
	return nil
}
