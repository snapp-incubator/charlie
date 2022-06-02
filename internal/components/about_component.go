package components

import (
	"syscall/js"

	oak "github.com/elliotforbes/go-webassembly-framework"
)

type AboutComponent struct{}

func init() {
	oak.RegisterFunction("do", Do)
}

func Do(this js.Value, inputs []js.Value) interface{} {
	println("Nice")

	return nil
}

func (a AboutComponent) Render() string {
	return `<div>
				<h2>My about component</h2>
				<button onClick="do();">Click me</button>
			</div>`
}
