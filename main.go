package main

import (
	"log"
)

func main() {
	img, err := loadImage("images/47b38cfdf64fa9d8.png")

	if err != nil {
		log.Fatal("Unable to load image!", err)
	}

	ctest := getCommonColors(img, 5)

	createCommonColorImage(hexListToColorList(ctest))

}
