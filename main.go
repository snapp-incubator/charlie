package main

import (
	"github.com/amirhnajafiz/fwfag/internal/components"
	"github.com/amirhnajafiz/fwfag/internal/router"
)

func main() {
	// creating a router
	router.NewRouter()
	router.RegisterRoute("home", components.Home)
	router.RegisterRoute("about", components.About)
}
