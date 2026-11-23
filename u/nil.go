package u

import (
	"reflect"
)

// Checks if nil
func Nil(x any) bool {
	if x == nil {
		return true
	}
	rv := reflect.ValueOf(x)
	// Checks typed nils (nil pointers, slices, maps, channels, funcs, interfaces).
	switch rv.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func, reflect.Interface:
		return rv.IsNil()
	default:
		return false
	}
}
