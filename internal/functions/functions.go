package functions

import "syscall/js"

func Register(name string, function func(value js.Value, i []js.Value) interface{}) {
	js.Global().Set(name, function)
}
