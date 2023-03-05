package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"image/color/palette"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			x1 := (float64(px) - float64(0.5)) / height * (xmax - xmin) + xmin
			y1 := (float64(py) - float64(0.5)) / width * (ymax - ymin) + ymin
			color1 := mandelbrot(complex(x1, y1))

			x2 := (float64(px) - float64(0.5)) / height * (xmax - xmin) + xmin
			y2 := (float64(py) + float64(0.5)) / width * (ymax - ymin) + ymin
			color2 := mandelbrot(complex(x2, y2))

			x3 := (float64(px) + float64(0.5)) / height * (xmax - xmin) + xmin
			y3 := (float64(py) - float64(0.5)) / width * (ymax - ymin) + ymin
			color3 := mandelbrot(complex(x3, y3))

			x4 := (float64(px) + float64(0.5)) / height * (xmax - xmin) + xmin
			y4 := (float64(py) + float64(0.5)) / width * (ymax - ymin) + ymin
			color4 := mandelbrot(complex(x4, y4))

			r1, g1, b1, _ := color1.RGBA()
			r2, g2, b2, _ := color2.RGBA()
			r3, g3, b3, _ := color3.RGBA()
			r4, g4, b4, _ := color4.RGBA()
			color := color.RGBA{uint8((r1 + r2 + r3 + r4)/4), uint8((g1 + g2 + g3 + g4)/4), uint8((b1 + b2 + b3 + b4)/4), 0xff}

			img.Set(px, py, color)
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 255
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette.Plan9[255-n]
		}
	}
	return color.Black
}
