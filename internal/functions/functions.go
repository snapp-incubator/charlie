package functions

import "syscall/js"

func Register(name string, function func(i []js.Value)) {
	js.Global().Set(name, function)
}
