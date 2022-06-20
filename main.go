package main

import (
	"fmt"
	"log"
)

func main() {
	img, err := loadImage("47b38cfdf64fa9d8.png")

	if err != nil {
		log.Fatal("Unable to load image!, err")
	}

	fmt.Println(getCommonColors(img, 5))
	fmt.Println(getCommonColors(img, 7))
}
