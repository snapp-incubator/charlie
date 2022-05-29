compile:
	GOOS=js GOARCH=wasm go build -o lib/lib.wasm main.go