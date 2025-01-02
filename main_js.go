package main

import (
	"fmt"
	"syscall/js"
)

func mdToHTMLJs(this js.Value, args []js.Value) any {
	// fmt.Sprintf("this=%+v, args=%+v", this, args)
	md := []byte(args[0].String())
	return string(mdToHTML(md))
}

// GOOS=js GOARCH=wasm go build -o markdown.wasm
func main() {
	fmt.Println("Hello, WebAssembly!")
	printMdToHTML()
	js.Global().Set("mdToHTML", js.FuncOf(mdToHTMLJs))
	select {}
}
