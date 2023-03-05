package main

import (
	"fmt"
	"math"
	"net/http"
	"log"
	"io"
	"strconv"
)

const (
	cells = 100
	xyrange = 30.0
	angle = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	color := r.Form["color"][0]
	if color == "" {
		color = "white"
	}

	heightString := "320"
	if len(r.Form["height"]) > 0 {
		heightString = r.Form["height"][0]
	}
	height, err := strconv.Atoi(heightString)
	if err != nil {
		log.Print(err)
	}

	widthString := "600"
	if len(r.Form["width"]) > 0 {
		widthString = r.Form["width"][0]
	}
	width, err := strconv.Atoi(widthString)
	if err != nil {
		log.Print(err)
	}

	xyscale := float64(width / 2 / xyrange)
	zscale := float64(height) * 0.4

	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, color, float64(width), float64(height), xyscale, zscale)
}

func surface(w io.Writer, color string, height float64, width float64, xyscale, zscale float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: %s; stroke-width: 0.7' " + 
		"width='%d' height='%d'>\n", color, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, height, xyscale, zscale)
			bx, by := corner(i, j, width, height, xyscale, zscale)
			cx, cy := corner(i, j+1, width, height, xyscale, zscale)
			dx, dy := corner(i+1, j+1, width, height, xyscale, zscale)
			fmt.Fprintf(w, "<polygon points='%g, %g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>\n")
}

func corner(i int, j int, width float64, height float64, xyscale float64, zscale float64) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := width / 2 + (x-y)*cos30*xyscale
	sy := height / 2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
