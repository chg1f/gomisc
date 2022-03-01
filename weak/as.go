package weak

import (
	"fmt"
	"reflect"
	"strconv"
)

func AsString(i interface{}) string {
	if i == nil {
		return ""
	} else if x, ok := i.(interface{ String() string }); ok {
		return x.String()
	}
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, v.Type().Bits())
	default:
		return fmt.Sprintf("%v", i)
	}
}
