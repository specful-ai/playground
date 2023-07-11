package main

import (
	"syscall/js"
)

func main() {
	document := js.Global().Get("document")
	canvas := document.Call("createElement", "canvas")
	document.Get("body").Call("appendChild", canvas)

	width := js.Global().Get("innerWidth").Int()
	height := js.Global().Get("innerHeight").Int()

	canvas.Set("width", width)
	canvas.Set("height", height)

	context := canvas.Call("getContext", "webgl")

	// Add your 3D rendering code here

	select {}
}
