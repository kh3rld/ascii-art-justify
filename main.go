package main

import (
	"fmt"
	"os"

	"ascii-art-justify/printart"
	"ascii-art-justify/reading"
)

func main() {
	var bannerFont string
	var alignFlag string
	var inputString string

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		return
	}

	switch len(args) {
	case 1:
		inputString = args[0]
	case 2:
		alignFlag = args[0]
		inputString = args[1]
	case 3:
		alignFlag = args[0]
		inputString = args[1]
		bannerFont = args[2]
	default:
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		return
	}

	switch bannerFont {
	case "standard":
		bannerFont = "standard.txt"
	case "shadow":
		bannerFont = "shadow.txt"
	case "thinkertoy":
		bannerFont = "thinkertoy.txt"
	default:
		bannerFont = "standard.txt"
	}

	bannerFile := reading.Reading(bannerFont)
	printart.PrintArt(bannerFile, inputString, alignFlag)
}
