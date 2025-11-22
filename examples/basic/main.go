package main

import (
	"bytes"
	_ "embed"
	"image/color"
	"log"

	"github.com/sago35/tinygo-image/png"
	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/examples/ili9341/initdisplay"
	"tinygo.org/x/drivers/ili9341"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	display := initdisplay.InitDisplay()
	display.SetRotation(ili9341.Rotation270)
	display.FillScreen(color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})

	for i := int16(0); i < 4; i++ {
		drawPngToDisplayer(display, 40*i, 20*i)
	}

	return nil
}

func drawPngToDisplayer(display drivers.Displayer, xxx, yyy int16) error {
	p := bytes.NewReader(pngImage)
	png.SetCallback(buffer[:], func(data []color.RGBA, x, y, w, h, width, height int16) {
		idx := 0
		for yy := y; yy < y+h; yy++ {
			for xx := x; xx < x+w; xx++ {
				if data[idx].A > 0 {
					display.SetPixel(xx+xxx, yy+yyy, data[idx])
				}
				idx++
			}
		}
	})

	_, err := png.Decode(p)
	return err
}

// Define the buffer required for the callback. In most cases, this setting
// should be sufficient.  For jpeg, the callback will always be called every
// 3*8*8*4 pix. png will be called every line, i.e. every width pix.
var buffer [3 * 8 * 8 * 4]color.RGBA

//go:embed tinygo-logo-alpha.png
var pngImage []byte
