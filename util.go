package goeasypost

import (
	"fmt"
	"reflect"
)

func flattenStructMap(start interface{}, prefix string, params map[string]string) error {

	s := reflect.Indirect(reflect.ValueOf(start))

	if s.Kind() == reflect.Slice {
		sl := reflect.ValueOf(start)
		for i := 0; i < sl.Len(); i++ {
			params[fmt.Sprintf("%v[%v]", prefix, i)] = fmt.Sprintf("%v", sl.Index(i).Interface())
		}

	} else {
		typeOfT := s.Type()
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)

			//skip blank values
			if !(isEmptyValue(reflect.ValueOf(f.Interface()))) {
				switch reflect.ValueOf(f.Interface()).Kind() {
				case reflect.Array, reflect.Map, reflect.Slice, reflect.Interface, reflect.Ptr:
					//potentially multi-valued types
					if "-" != typeOfT.Field(i).Tag.Get("json") { //skip json-designated noop fields

						err := flattenStructMap(f.Interface(), fmt.Sprintf("%v[%v]", prefix, typeOfT.Field(i).Tag.Get("json")), params)
						if err != nil {
							return err
						}
					}
				default:
					if "-" != typeOfT.Field(i).Tag.Get("json") { //skip json-designated noop fields
						params[fmt.Sprintf("%v[%v]", prefix, typeOfT.Field(i).Tag.Get("json"))] = fmt.Sprintf("%v", f.Interface())
					}
				}

			}
		}
	}
	return nil
}

func isEmptyValue(v reflect.Value) bool {

	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
