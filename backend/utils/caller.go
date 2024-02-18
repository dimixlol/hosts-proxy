package utils

import (
	"reflect"
)

type Method struct {
	MethodName string
	Arguments  []any
}

// Caller calls a method of an object with the given input
func Caller(o interface{}, method *Method) {
	m := reflect.ValueOf(o).MethodByName(method.MethodName)
	in := make([]reflect.Value, m.Type().NumIn())
	for i := 0; i < m.Type().NumIn(); i++ {
		in[i] = reflect.ValueOf(method.Arguments[i])
	}
	m.Call(in)
}
