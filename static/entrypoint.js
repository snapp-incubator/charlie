// framework name
const frameworkName = "./../lib/lib.wasm";

// creating a new go instance to create oir framework in lib.was
const go = new Go();

WebAssembly.instantiateStreaming(fetch(frameworkName), go.importObject)
    .then(result => {
        go.run(result.instance)
            .then(r => console.log(r))
            .catch(e => {
                console.error(e)
                console.log("Error for go run")
            })
    })
    .catch(e => {
        console.error(e)
    })
