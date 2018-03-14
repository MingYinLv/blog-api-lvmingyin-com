package util

import (
	"reflect"
	"strings"
)

func MapSet(obj interface{}, maps map[string]interface{}) interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(&obj).Elem()

	for k := 0; k < t.NumField(); k++ {
		name := t.Field(k).Name
		lowerName := strings.ToLower(name)
		val, ok := maps[lowerName]
		if lowerName == "id" {
			val = int64(val.(int))
		}
		if ok {
			v.FieldByName(name).Set(reflect.ValueOf(val))
		}
	}

	return obj
}
