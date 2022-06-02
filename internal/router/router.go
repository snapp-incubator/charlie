package router

import "github.com/amirhnajafiz/fwfag/internal/components"

type Router struct {
	Routes map[string]components.Component
}

var router Router

func init() {
	router.Routes = make(map[string]components.Component)
}
