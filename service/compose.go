package service

import (
	"reflect"
)

// Compose applies middlewares to Service.
func Compose(s interface{}, mws ...interface{}) interface{} {
	for i := len(mws) - 1; i >= 0; i-- {
		vv := reflect.ValueOf(mws[i]).Call([]reflect.Value{reflect.ValueOf(s)})
		s = vv[0].Interface()
	}
	return s
}
