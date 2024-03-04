package image

import (
	"bytes"
	"image"
	"syscall/js"
)

func (i *Image) setupOnImgLoad() {
	i.onImgLoadCb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		array := args[0]
		i.inBuf = make([]uint8, array.Get("bytelength").Int())
		js.CopyBytesToGo(i.inBuf, array)

		reader := bytes.NewReader(i.inBuf)
		var err error
		i.sourceImg, _, err = image.Decode(reader)
		if err != nil {
			i.log(err.Error())
			return nil
		}
		i.log("Ready for operations")
		// reset brightness and contrast sliders
		js.Global().Get("document").
			Call("getElementById", "brightness").
			Set("value", 0)

		js.Global().Get("document").
			Call("getElementById", "contrast").
			Set("value", 0)
		return nil
	})
}
