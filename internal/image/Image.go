package image

import (
	"bytes"
	"image"
	"syscall/js"
)

type Image struct {
	inBuf                    []uint8
	outBuf                   bytes.Buffer
	onImgLoadCb, shutdownCb  js.Func
	brightnessCb, contrastCb js.Func
	hueCb, satCb             js.Func
	sourceImg                image.Image

	console js.Value
	done    chan struct{}
}

func New() *Image {
	return &Image{
		console: js.Global().Get("console"),
		done:    make(chan struct{}),
	}
}
func (i *Image) Start() {
	//Setup Callbacks
	i.setupOnImgLoad()
	i.clog("This is me.")
	// js.Global().Set("loadImage", i.onImgLoadCb)

	//shutdown listner
	i.setupShutdownCb()
	js.Global().Get("document").
		Call("getElementById", "close").
		Call("addEventListener", "click", i.shutdownCb)

		//when shutdown
	<-i.done
	i.log("Shutting down app")
	i.onImgLoadCb.Release()
	i.brightnessCb.Release()
	i.contrastCb.Release()
	i.hueCb.Release()
	i.satCb.Release()
	i.shutdownCb.Release()
}
func (i *Image) clog(msg string) {
	js.Global().Get("console").
		Call("log", msg)
}
func (i *Image) log(msg string) {
	js.Global().Get("document").
		Call("getElementById", "status").
		Set("innerText", msg)
}
func (i *Image) setupShutdownCb() {
	i.shutdownCb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		i.done <- struct{}{}
		return nil
	})
}
