package db

import (
	"fmt"
	"reflect"
	"strings"
)

func objToString(obj interface{}) string {
	objT := reflect.TypeOf(obj)
	var fields []string
	for i := 0; i < objT.NumField(); i++ {
		fieldT := objT.Field(i)
		tag := fieldT.Tag.Get("db")
		if tag == "" {
			continue
		}
		field := fmt.Sprintf("`%s`", tag)
		fields = append(fields, field)
	}
	return strings.Join(fields, ",")
}
