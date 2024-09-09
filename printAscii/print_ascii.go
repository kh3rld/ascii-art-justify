package output

import (
	"fmt"
	"strings"
)

func PrintArt(str string, asciiArtGrid [][]string, align string) error {
	switch str {
	case "":
		fmt.Print()
	case "\\n":
		fmt.Println()
	case "\\r", "\\f", "\\v", "\\t", "\\b", "\\a":
		return fmt.Errorf("error: unsupported escape sequence '%s'", str)
	default:
		s := strings.ReplaceAll(str, "\\n", "\n")
		s = strings.ReplaceAll(s, "\\r", "\r")
		s = strings.ReplaceAll(s, "\\f", "\f")
		s = strings.ReplaceAll(s, "\\v", "\v")
		s = strings.ReplaceAll(s, "\\t", "\t")
		s = strings.ReplaceAll(s, "\\b", "\b")
		s = strings.ReplaceAll(s, "\\a", "\a")
		words := strings.Split(s, "\n")
		num := 0
		for _, word := range words {
			if word == "" {
				num++
				if num < len(words) {
					fmt.Println()
					continue
				}
			} else {
				// Get the ASCII art for the word
				var artLines []string
				for i := 1; i <= 8; i++ {
					var line string
					for _, char := range word {
						index := int(char - 32)
						if index < 0 || index >= len(asciiArtGrid) {
							return fmt.Errorf("unknown character: %q", char)
						} else {
							line += asciiArtGrid[index][i]
						}
					}
					artLines = append(artLines, line)
				}

			}
		}
	}
	return nil
}
