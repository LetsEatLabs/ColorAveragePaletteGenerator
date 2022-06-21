package main

// Adapted from https://golangcode.com/find-common-colours-in-image/
// Edited to support PNG files, among other things

import (
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/EdlinOrg/prominentcolor"
)

func getCommonColorsMean(img image.Image, k int) []string {

	cols, err := prominentcolor.KmeansWithAll(k, img, prominentcolor.ArgumentAverageMean|prominentcolor.ArgumentNoCropping,
		255,
		prominentcolor.GetDefaultMasks())

	if err != nil {
		log.Fatal("Failed to process image", err)
	}

	var dom_cols []string

	for _, colour := range cols {
		dom_cols = append(dom_cols, fmt.Sprintf("#%s", colour.AsString()))
	}

	return dom_cols
}

func getCommonColorsMedian(img image.Image, k int) []string {

	cols, err := prominentcolor.KmeansWithAll(k, img, prominentcolor.ArgumentNoCropping,
		255,
		prominentcolor.GetDefaultMasks())

	if err != nil {
		log.Fatal("Failed to process image", err)
	}

	var dom_cols []string

	for _, colour := range cols {
		dom_cols = append(dom_cols, fmt.Sprintf("#%s", colour.AsString()))
	}

	return dom_cols
}

func getAllTargetImages(dirpath string) []string {
	var targetimages []string

	files, err := ioutil.ReadDir("./images")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		targetimages = append(targetimages, f.Name())
	}

	return targetimages
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
