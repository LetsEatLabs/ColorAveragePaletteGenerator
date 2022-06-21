package main

import (
	"fmt"
	"log"
)

// Consts
const SAVE_DIR = "./averages/"
const LOAD_DIR = "./images/"

func main() {

	targetimages := getAllTargetImages(LOAD_DIR)

	for im := range targetimages {

		img, err := loadImage(fmt.Sprintf("%s%s", LOAD_DIR, targetimages[im]))

		if err != nil {
			log.Fatal("Unable to load image!", err)
		}
		meancols := getCommonColorsMean(img, 5)
		createCommonColorImage(hexListToColorList(meancols), fmt.Sprintf("%s%s.average.png", SAVE_DIR, targetimages[im]))
	}

}
