package weak

import (
	"errors"
	"reflect"
)

var ErrExpectMap = errors.New("expect map")
var ErrExpectPtr = errors.New("expect ptr")

type Merger struct {
	IgnoreIncrease   bool
	IgnoreEmpty      bool
	OverwriteDiscord bool
}

var standardMerger = &Merger{}

func Merge(i, o interface{}) error {
	return standardMerger.Merge(i, o)
}
func (m *Merger) Merge(i, o interface{}) error {
	if it := reflect.TypeOf(i); it.Kind() != reflect.Map {
		return ErrExpectMap
	} else if ot := reflect.TypeOf(o); ot.Kind() != reflect.Ptr {
		return ErrExpectPtr
	} else if ot.Elem().Kind() != reflect.Map {
		return ErrExpectMap
	}
	iv := reflect.ValueOf(i)
	ov := reflect.ValueOf(o)
	for _, ik := range iv.MapKeys() {
		icv := iv.MapIndex(ik).Elem()
		if m.IgnoreEmpty && isEmpty(icv) {
			continue
		}
		if !ov.Elem().MapIndex(ik).IsValid() {
			if !m.IgnoreIncrease {
				ov.Elem().SetMapIndex(ik, icv)
			}
			continue
		}
		ocv := ov.Elem().MapIndex(ik).Elem()
		if icv.Type().Kind() != ocv.Type().Kind() {
			if m.OverwriteDiscord {
				ov.Elem().SetMapIndex(ik, icv)
			}
			continue
		}
		ov.Elem().SetMapIndex(ik, icv)
	}
	return nil
}
