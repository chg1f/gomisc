package weak

import (
	"errors"
	"fmt"
	"strings"
)

type WeakError struct {
	I interface{}
}

func (e WeakError) Error() string {
	if ret, ok := e.I.(interface {
		String() string
	}); ok {
		return ret.String()
	}
	return fmt.Sprintf("%v", e.I)
}
func (e WeakError) As(interface{}) bool {
	return false
}
func (e WeakError) Is(error) bool {
	return false
}

func AsError(i interface{}) error {
	return WeakError{I: i}
}

type CompressError []error

func (es CompressError) Error() string {
	t := make([]string, 0, len(es))
	for _, e := range es {
		t = append(t, e.Error())
	}
	return strings.Join(t, "; ")
}
func (es CompressError) As(t interface{}) bool {
	for _, e := range es {
		if errors.As(e, t) {
			return true
		}
	}
	return false
}
func (es CompressError) Is(t error) bool {
	for _, e := range es {
		if errors.Is(e, t) {
			return true
		}
	}
	return false
}

// TODO:
func (es CompressError) Shrink() CompressError {
	if len(es) == 0 {
		return nil
	}
	return es
}
