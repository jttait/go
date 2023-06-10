package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"image/gif"
	"io"
	"os"
	"flag"
)

func main() {
	var outputFlag = flag.String("output", "jpeg", "Output format")
	flag.Parse()

	if *outputFlag == "jpeg" {
		if err := toJPEG(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
			os.Exit(1)
		}
	} else if *outputFlag == "png" {
		if err := toPNG(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "png: %v\n", err)
			os.Exit(1)
		}
	} else if *outputFlag == "gif" {
		if err := toGIF(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "png: %v\n", err)
			os.Exit(1)
		}
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format = ", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toPNG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err !=nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format = ", kind)
	return png.Encode(out, img)
}

func toGIF(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err !=nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format = ", kind)
	return gif.Encode(out, img, &gif.Options{})
}
