package weak

import (
	"fmt"
	"reflect"
	"strconv"
)

func AsString(value interface{}) string {
	if value == nil {
		return ""
	}
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.String {
		return v.String()
	} else if v.Type() == reflect.TypeOf([]byte("")) {
		return string(v.Bytes())
	} else if x, ok := value.(interface{ String() string }); ok {
		return x.String()
	}
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", value)
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", value)
	default:
		return fmt.Sprintf("%v", value)
	}
}

func AsNumber(value interface{}) float64 {
	if value == nil {
		return 0
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Float32, reflect.Float64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Float()
	}
	n, err := strconv.ParseFloat(weak.AsString(value), 64)
	if err != nil {
		return 0
	}
	return n
}
