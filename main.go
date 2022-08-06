package main

import (
	"github.com/amirhnajafiz/assembled/internal/components"
	"github.com/amirhnajafiz/assembled/internal/router"
)

func main() {
	// creating a router
	router.NewRouter()
	router.RegisterRoute("home", components.Home)
	router.RegisterRoute("about", components.About)
}
