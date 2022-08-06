# Assembled

<p align="center">
	<img src=".github/readme/logo.jpeg" width="500" alt="logo" /><br />
	Writing a Frontend Web Framework with WebAssembly And Go.
</p>

---

In this project I build the base of an incredibly simple frontend 
framework written in Go that compiles into WebAssembly. 

### Wasm and JS
JavaScript has been the only programming language that the browser 
understands. JavaScript has stood the test of time and it has been 
able to deliver the performance needed by most web applications. 
But when it comes to 3D games, VR, AR, and image editing apps, 
JavaScript is not quite up to the mark since it is interpreted. 
Although JavaScript engines such as Gecko and V8 have Just in Time 
compilation capabilities, JavaScript is not able to provide the high 
performance required by modern web applications.

WebAssembly(also known as wasm) is meant to solve this problem. 
WebAssembly is a virtual assembly language for the browser. 
When we say virtual, it means that it cannot be run natively 
on the underlying hardware. Since the browser can be running on any 
architecture, it is not possible for the browser to run WebAssembly 
directly on the underlying hardware. But this highly optimized virtual 
assembly format can be processed much quicker than vanilla JavaScript 
by modern browsers since it is compiled and is more close to the 
hardware architecture than JavaScript. The following figure shows 
where WebAssembly stands in the stack when compared to Javascript. 
It is closer to the Hardware than JavaScript.

### Project
We will create a simple application that is used to format JSON. 
If a JSON without any formatting is passed as input, 
it will be formatted and printed.

#### Example
Input:
```json
{"website":"golangbot.com", "tutorials": {"string":"https://golangbot.com/strings/", "maps":"https://golangbot.com/maps/", "goroutine":"https://golangbot.com/goroutines/", "channels":"https://golangbot.com/channels/"}}
```

Output:
```json
{
  "tutorials": {
    "channels": "https://golangbot.com/channels/",
    "goroutine": "https://golangbot.com/goroutines/",
    "maps": "https://golangbot.com/maps/",
    "string": "https://golangbot.com/strings/"
  },
  "website": "golangbot.com"
}
```

### Structure
```shell
assembled
    ├── assets
        └── index.html
        └── json.wasm
        └── wasm_exec.js
    └── cmd
        ├── server
          └── main.go
        └── wasm
          └── main.go
```

### How to run?
```shell
go run ./cmd/server/main.go
```

Checkout **localhost:9090**.
