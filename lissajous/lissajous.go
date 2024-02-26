package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.White,
	color.RGBA{255, 101, 80, 190},
	color.RGBA{250, 230, 80, 255},
	color.RGBA{162, 192, 59, 200},
	color.RGBA{162, 192, 59, 200},
}

const (
	whiteIndex = 0
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nFrames = 64
		delay   = 8
	)
	freq := rand.Float64() * 3
	anim := gif.GIF{LoopCount: nFrames}

	phase := 0.1
	minIndex := 2
	maxIndex := len(palette)
	colorIndex := uint8(rand.Intn(maxIndex-minIndex+1) + minIndex)

	for i := 0; i < nFrames; i++ {

		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

	}
	gif.EncodeAll(out, &anim)
}
