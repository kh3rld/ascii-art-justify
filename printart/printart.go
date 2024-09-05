package printart

import (
	"fmt"
	"strings"
)

func PrintArt(bannerFileSlice []string, inputString string, alignFlag string) {
	align := ""
	if alignFlag != "" {
		align = strings.ToLower(alignFlag[8:])

	} else {
		align = "left"
		// fmt.Println("align should have a value")
	}

	if inputString == "\\n" {
		fmt.Println()
		return
	} else if inputString == "" {
		return
	} else if inputString == "\\t" {
		fmt.Println("	")
		return
	}

	// Handle unprintable sequences
	unprintableSequences := []string{"\\a", "\\b", "\\v", "\\f", "\\r"}

	for _, unprintable := range unprintableSequences {
		if strings.Contains(inputString, unprintable) {
			fmt.Println("Input string contains an unprintable sequence")
			return
		}
	}

	tabCharText := strings.Replace(inputString, "\\t", "    ", -1)
	newlineCharText := strings.ReplaceAll(tabCharText, "\\n", "\n")
	splitArguments := strings.Split(newlineCharText, "\n")

	// Handle foreign inputs
	for _, splitArg := range splitArguments {
		for _, char := range splitArg {
			if char < 32 || char > 126 {
				fmt.Println("Input string contains unprintable character")
				return
			}
		}
	}

	for _, text := range splitArguments {
		if text == "" {
			fmt.Println()
			continue
		}

		const asciiHeight = 8
		// fmt.Println(align)

		for j := 0; j < asciiHeight; j++ {
			switch align {
			case "left":
				for _, char := range text {

					startingIndex := int(char-32)*9 + 1
					fmt.Printf(bannerFileSlice[startingIndex+j])
				}
				fmt.Println()
			case "rignt":
				for _, char := range text {

					startingIndex := int(char-32)*9 + 1
					fmt.Printf(bannerFileSlice[startingIndex+j])
				}
				fmt.Println()
			case "center":
				for _, char := range text {

					startingIndex := int(char-32)*9 + 1
					fmt.Printf(bannerFileSlice[startingIndex+j])
				}
				fmt.Println()
			case "justify":
				for _, char := range text {

					startingIndex := int(char-32)*9 + 1
					fmt.Printf(bannerFileSlice[startingIndex+j])
				}
				fmt.Println()
			}
		}
	}
}
