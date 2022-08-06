package router

import (
	"syscall/js"

	"github.com/amirhnajafiz/assembled/internal/components"
)

type Router struct {
	Routes map[string]components.Component
}

var router Router

func init() {
	router.Routes = make(map[string]components.Component)
}

func NewRouter() {
	js.Global().Set("Link", Link)
	js.Global().Get("document").Call("getElementById", "view").Set("innerHTML", "")
}

func RegisterRoute(path string, component components.Component) {
	router.Routes[path] = component
}

func Link(value js.Value, i []js.Value) interface{} {
	println("Link Hit")

	comp := router.Routes[i[0].String()]
	html := comp.Render()

	js.Global().Get("document").Call("getElementById", "view").Set("innerHTML", html)

	return nil
}
