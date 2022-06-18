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

func main() {

	// Step 1: Load the image
	img, err := loadImage(os.Args[1])
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to load image, %s\n", os.Args[1]), err)
	}

	// Step 2: Process it
	colours, err := prominentcolor.Kmeans(img)
	if err != nil {
		log.Fatal("Failed to process image", err)
	}

	fmt.Println("Dominant colours:")
	for _, colour := range colours {
		fmt.Println("#" + colour.AsString())
	}
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
