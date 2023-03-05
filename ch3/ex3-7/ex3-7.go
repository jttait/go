package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py) / height * (ymax - ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px) / width * (xmax - xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newtons(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func newtons(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var f complex128
	var derivative complex128
	for n := uint8(0); n < iterations; n++ {
		f = z*z*z*z - 1
		derivative = 4*z*z*z
		if cmplx.Abs(f) < 0.001 {
			return color.Gray{255 - contrast*n}
		}
		z -= f / derivative 
	}
	return color.Black
}
