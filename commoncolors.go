package main

// Adapted from https://golangcode.com/find-common-colours-in-image/
// Edited to support PNG files, among other things

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/EdlinOrg/prominentcolor"
)

func getCommonColors(img image.Image, k int) []string {

	colours, err := prominentcolor.KmeansWithAll(k, img, prominentcolor.ArgumentAverageMean|prominentcolor.ArgumentNoCropping,
		prominentcolor.DefaultSize,
		prominentcolor.GetDefaultMasks())

	if err != nil {
		log.Fatal("Failed to process image", err)
	}

	var dom_colours []string

	for _, colour := range colours {
		dom_colours = append(dom_colours, fmt.Sprintf("#%s", colour.AsString()))
	}

	return dom_colours
}

func loadImage(fileInput string) (image.Image, error) {
	f, err := os.Open(fileInput)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, err := png.Decode(f)
	return img, err
}
