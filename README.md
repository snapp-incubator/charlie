# FWFAG

Writing a Frontend Web Framework with WebAssembly And Go.

In this project I build the base of an incredibly simple frontend 
framework written in Go that compiles into WebAssembly. 

At a minimum, this will include the following features:
- Function Registration
- Components
- Routing

## Static
Our entrypoint is the main script which
will generate our goland framework from lib directory.

## lib
Lib will be our framework output file that
the entrypoint will get and build the framework.

## Functions
To register our functions, we use 
javascript syscall in golang.
```go
import (
	"syscall/js"
)

func Register(name string, function func(i []js.Value)) {
	js.Global().Set(name, function)
}
```

## Components
React tends to feature classes that feature 
_render()_ functions which return the HTML/JSX/whatever
code you wish to render for said component. 

Our components will do the same:
```go
package components

type Component interface {
	Render() string
}
```
