package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	colors "gopkg.in/go-playground/colors.v1"
)

func hexListToColorList(cols []string) []*colors.RGBColor {
	var rgbs []*colors.RGBColor

	for c := range cols {
		fmtc, err := colors.ParseHEX(cols[c])

		if err != nil {
			log.Fatal("Could not parse color ", err)
		}
		rgbs = append(rgbs, fmtc.ToRGB())
	}

	return rgbs
}

func createCommonColorImage(cols []*colors.RGBColor) {

	height := 500
	width := 200
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Set color for each pixel.
	curr_color := color.RGBA{cols[0].R, cols[0].G, cols[0].B, 255}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if y%100 == 0 {
				curr_color = color.RGBA{cols[y/100].R, cols[y/100].G, cols[y/100].B, 255}
			}

			img.Set(x, y, curr_color)
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
