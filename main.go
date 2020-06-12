package main

import (
	"strings"
	"syscall/js"

	"go.starlark.net/starlark"
)

func main() {
	js.Global().Set("evalStarlark", js.FuncOf(eval))

	// Prevent the app from terminating.
	c := make(chan struct{})
	<-c
}

func toJson(d starlark.StringDict) string {
	buf := new(strings.Builder)
	buf.WriteByte('{')
	sep := ""
	for _, name := range d.Keys() {
		if strings.HasPrefix(name, "_") {
			continue // skip private symbols
		}
		if _, ok := d[name].(starlark.Callable); ok {
			continue // skip functions
		}
		buf.WriteString(sep)
		buf.WriteString("\"")
		buf.WriteString(name)
		buf.WriteString("\"")
		buf.WriteString(": ")
		buf.WriteString(d[name].String())
		sep = ", "
	}
	buf.WriteByte('}')
	return buf.String()
}

func printer(thread *starlark.Thread, msg string) {
	js.Global().Call("updateOutput", msg)
}

func eval(this js.Value, inputs []js.Value) interface{} {
	text := inputs[0].String()
	thread := &starlark.Thread{
		Name:  "main",
		Print: printer,
	}
	res, err := starlark.ExecFile(thread, "", text, nil)
	if err == nil {
		js.Global().Call("updateJSON", toJson(res))
	} else {
		js.Global().Call("setError", err.Error())
	}

	return nil
}
