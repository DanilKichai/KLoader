package common

import "reflect"

func New[T any]() T {
	var t T
	theType := reflect.TypeOf(t)

	var newValue reflect.Value
	if theType.Kind() == reflect.Pointer {
		newValue = reflect.New(theType.Elem())
	} else {
		newValue = reflect.New(theType).Elem()
	}

	return newValue.Interface().(T)
}

func Nil[T any]() (nilT T) {
	return nilT
}
