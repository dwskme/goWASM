package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello World!")
	document := js.Global().Get("document")
	p := document.Call("createElement", "p")
	p.Set("innerHTML", "Hello WASM from Go!")
	p.Set("className", "block")

	styles := document.Call("createElement", "style")
	styles.Set("innerHTML", `
		.block {
			border: 1px solid black; color: white; background: black;
		}
	`)

	document.Get("head").Call("appendChild", styles)
	document.Get("body").Call("appendChild", p)
}
