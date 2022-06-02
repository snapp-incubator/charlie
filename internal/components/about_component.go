package components

import (
	"syscall/js"

	"github.com/amirhnajafiz/fwfag/internal/functions"
)

type AboutComponent struct{}

var About AboutComponent

func init() {
	functions.Register("do", Do)
}

func Do(_ js.Value, _ []js.Value) interface{} {
	println("Nice")

	return nil
}

func (a AboutComponent) Render() string {
	return `<div>
				<h2>My about component</h2>
				<button onClick="do();">Click me</button>
			</div>`
}
