// Copyright 2017 Marin Basic <marin@marin-basic.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package gomercedesdealer

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

// covert struct fields to map of strings
func mapFromStruct(i interface{}) map[string]string {
	params := make(map[string]string)
	s := reflect.ValueOf(i).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		var value string
		var key string = typeOfT.Field(i).Name

		switch s.Field(i).Kind().String() {
		case "slice":
			value = strings.Join(s.Field(i).Interface().([]string), ",")
			break
		case "float64":
			if s.Field(i).Interface().(float64) <= 0 {
				break
			}
			value = fmt.Sprintf("%f", s.Field(i).Interface().(float64))
			break
		case "int":
			if s.Field(i).Interface().(int) <= 0 {
				break
			}
			value = fmt.Sprintf("%d", s.Field(i).Interface().(int))
			break
		default:
			value = s.Field(i).Interface().(string)
		}

		if len(value) > 0 {
			params[camelCase(key)] = value
		}

	}
	return params
}

// first char to lower
func camelCase(s string) string {
	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}
