package main

import (
	"log"

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

func createCommonColorImage(cols []colors.RGBColor, height int, width int) bool {
	return true
}
