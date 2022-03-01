package weak

import (
	"reflect"
	"strconv"
	"strings"
)

type Flater struct {
	StructTag    string
	UintptrValue bool
	FlatMap      bool
	FlatStruct   bool
	FlatArray    bool
	FlatSlice    bool
}

var standareFlater = &Flater{
	StructTag:    "",
	UintptrValue: true,
	FlatMap:      true,
	FlatStruct:   true,
	FlatArray:    true,
	FlatSlice:    true,
}

func (f *Flater) Flat(input interface{}) (output map[string]interface{}) {
	return f.FlatValue(reflect.ValueOf(input)).Interface().(map[string]interface{})
}
func (f *Flater) FlatValue(input reflect.Value) (output reflect.Value) {
	output = reflect.ValueOf(make(map[string]interface{}))
	input = Unwrap(input)
	switch input.Kind() {
	case reflect.Array, reflect.Slice:
		if !f.FlatArray && input.Kind() == reflect.Array {
			goto setIn
		} else if !f.FlatSlice && input.Kind() == reflect.Slice {
			goto setIn
		}
		for index := 0; input.Len() > index; index += 1 {
			child := f.FlatValue(input.Index(index))
			for _, key := range child.MapKeys() {
				value := Unwrap(child.MapIndex(key))
				if f.UintptrValue && value.CanAddr() {
					value = value.Addr()
				}
				output.SetMapIndex(
					reflect.ValueOf("["+strconv.Itoa(index)+"]"+key.String()),
					value,
				)
			}
		}
		return
	case reflect.Map:
		if !f.FlatMap {
			goto setIn
		}
		for _, index := range input.MapKeys() {
			child := f.FlatValue(input.MapIndex(index))
			index = Unwrap(index)
			for _, key := range child.MapKeys() {
				value := Unwrap(child.MapIndex(key))
				if f.UintptrValue && value.CanAddr() {
					value = value.Addr()
				}
				if key.String() != "" {
					output.SetMapIndex(
						reflect.ValueOf(AsString(index.Interface())+"."+key.String()),
						value,
					)
				} else {
					output.SetMapIndex(
						reflect.ValueOf(AsString(index.Interface())),
						value,
					)
				}
			}
		}
		return
	case reflect.Struct:
		if !f.FlatStruct {
			goto setIn
		}
		for index := 0; input.NumField() > index; index += 1 {
			child := f.FlatValue(input.Field(index))
			name := input.Type().Field(index).Name
			tag := input.Type().Field(index).Tag.Get(f.StructTag)
			if parts := strings.SplitN(tag, ",", 2); len(parts) >= 1 && parts[0] != "" {
				name = parts[0]
			}
			for _, key := range child.MapKeys() {
				value := Unwrap(child.MapIndex(key))
				if f.UintptrValue && value.CanAddr() {
					value = value.Addr()
				}
				if key.String() != "" {
					output.SetMapIndex(
						reflect.ValueOf(name+"."+key.String()),
						value,
					)
				} else {
					output.SetMapIndex(
						reflect.ValueOf(name),
						value,
					)
				}
			}
		}
		return
	}
setIn:
	output.SetMapIndex(
		reflect.ValueOf(""),
		input,
	)
	return
}

func Unwrap(v reflect.Value) reflect.Value {
	switch v.Kind() {
	case reflect.Interface, reflect.Uintptr:
		return Unwrap(v.Elem())
	}
	return v
}
