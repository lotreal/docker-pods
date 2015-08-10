package sh

import (
	"fmt"
	"reflect"
)

// func InterfaceSlice(in inteface{}) []interface{} {
// 	s := reflect.ValueOf(in)
// 	if s.Kind() != reflect.Slice {
// 		panic("InterfaceSlice() given a non-slice type")
// 	}

// 	ret := make([]interface{}, s.Len())

// 	for i:=0; i<s.Len(); i++ {
// 		ret[i] = s.Index(i).Interface()
// 	}

// 	return ret
// }

func TabWrite(v interface{}) {

	s := reflect.ValueOf(v)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i:=0; i<s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	fmt.Printf("%#v", ret)
	// for i, val := range v {
	// 	fmt.Printf("%s", val)

	// }
	// s  := reflect.ValueOf(v)
	// st := reflect.TypeOf(v)

	// fmt.Printf("%s", v)
	// fmt.Printf("value: %#v", s)
	// fmt.Printf("type: %#v", st)
	// record, err := reader.Read()
	// if err != nil {
	// 	return err
	// }
	// if record == "" {
	// 	return errors.New("record is empty")
	// }

	// s := reflect.ValueOf(v).Elem()
	// for i := 0; i < s.NumField(); i++ {
	// 	f := s.Field(i)

	// 	field := reader.fields[i]
	// 	start := field.Start
	// 	end   := field.End

	// 	var value string
	// 	if end == 0 {
	// 		value = record[start:]
	// 	} else {
	// 		value = record[start:end]
	// 	}

	// 	f.SetString(strings.Trim(value, " "))
	// }
	// return nil
}
