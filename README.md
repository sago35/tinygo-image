# tinygo-image

This package is maintained to make Go’s image package usable with TinyGo.

It’s basically the same as tinygo.org/x/drivers/image, but the data type used in the callback has been changed from RGB565 (uint16) to color.RGBA.
This allows branching based on whether a pixel is transparent.
On the other hand, the data size becomes larger.
