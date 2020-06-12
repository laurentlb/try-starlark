# Try Starlark

Try Starlark is an online Starlark interpreter. Play with the language from your browser.

[Starlark](https://github.com/bazelbuild/starlark) is a simple language that
looks like Python. It is designed to be embedded in other applications.

This repository uses [starlark-go](https://github.com/google/starlark-go), an
implementation of a Starlark interpreter written in Go. It is compiled to
WebAssembly and used from Javascript.


## Compile

```
$ cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
$ GOOS=js GOARCH=wasm go build
```
