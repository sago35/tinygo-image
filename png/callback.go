package png

import (
	"image/color"
)

var (
	callback    Callback = func(data []color.RGBA, x, y, w, h, width, height int16) {}
	callbackBuf []color.RGBA
)

// A portion of the image data consisting of data, x, y, w, and h is passed to
// Callback. The size of the whole image is passed as width and height.
type Callback func(data []color.RGBA, x, y, w, h, width, height int16)

// SetCallback registers the buffer and fn required for Callback. Callback can
// be called multiple times by calling Decode().
func SetCallback(buf []color.RGBA, fn Callback) {
	callbackBuf = buf
	callback = fn
}
