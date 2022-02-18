package weak

import "reflect"

func isEmpty(iv reflect.Value) bool {
	if !iv.IsValid() {
		return true
	}
	switch iv.Type().Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return iv.Len() == 0
	case reflect.Interface, reflect.Ptr:
		return iv.IsNil()
	default:
		return iv.IsZero()
	}
}
func IsEmpty(i interface{}) bool {
	return isEmpty(reflect.ValueOf(i))
}
