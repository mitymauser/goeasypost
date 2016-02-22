/*
 * Copyright (c) 2016 Stewart Buskirk <mitymauser@gmail.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */
package goeasypost

import (
	"fmt"
	"reflect"
)

func flattenStructMap(start interface{}, prefix string, params map[string]string) error {

	s := reflect.Indirect(reflect.ValueOf(start))

	//todo other types, but not relevant to easypost...
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
				case reflect.Array, reflect.Map, reflect.Slice, reflect.Interface, reflect.Ptr, reflect.Struct:
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
