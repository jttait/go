package main

import (
	"image"
	"image/color"
	"math/cmplx"
	//"image/png"
	//"os"
	"time"
	"fmt"
	"sync"
)

type pixel struct {
	px int
	py int
	color color.Color
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	/*start := time.Now()
	for py := 0; py < height; py++ {
		y := float64(py) / height * (ymax - ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px) / width * (xmax - xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	fmt.Println(time.Since(start))*/

	start := time.Now()
	pixels := make(chan pixel, height*width)
	var wg sync.WaitGroup
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			wg.Add(1)
			go func(px int, py int) {
				defer wg.Done()
				y := float64(py) / height * (ymax - ymin) + ymin
				x := float64(px) / width * (xmax - xmin) + xmin
				z := complex(x, y)
				pixels <- pixel{px, py, mandelbrot(z)}
			}(px, py)
		}
	}

	go func() {
		wg.Wait()
		close(pixels)
	}()

	for pixel := range pixels {
		img.Set(pixel.px, pixel.py, pixel.color)
	}
	//png.Encode(os.Stdout, img)
	fmt.Println(time.Since(start))
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
