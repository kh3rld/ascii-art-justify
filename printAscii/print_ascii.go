package justify

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

func PrintArt(str string, asciiArtGrid [][]string, align string) error {
	terminalWidth, _, err := getTerminalSize() // Function to get terminal width
	if err != nil {
		return err
	}

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

				// Align the ASCII art
				alignedArt := alignArt(artLines, align, terminalWidth)
				fmt.Println(alignedArt)
			}
		}
	}
	return nil
}

func alignArt(artLines []string, align string, width int) string {
	var alignedLines []string
	for _, line := range artLines {
		switch align {
		case "center":
			alignedLines = append(alignedLines, centerAlign(line, width))
		case "left":
			alignedLines = append(alignedLines, line)
		case "right":
			alignedLines = append(alignedLines, rightAlign(line, width))
		case "justify":
			alignedLines = append(alignedLines, justifyAlign(line, width))
		}
	}
	return strings.Join(alignedLines, "\n")
}

func centerAlign(line string, width int) string {
	padding := (width - len(line)) / 2
	return strings.Repeat(" ", padding) + line
}

func rightAlign(line string, width int) string {
	padding := width - len(line)
	return strings.Repeat(" ", padding) + line
}

func justifyAlign(line string, width int) string {
	// If the line is longer than or equal to the width, return the line as is
	if len(line) >= width {
		return line
	}

	// Split the line into words
	words := strings.Fields(line)
	if len(words) == 1 {
		// If there's only one word, pad it to the right
		return line + strings.Repeat(" ", width-len(line))
	}

	// Calculate total spaces needed
	totalSpaces := width - len(line)
	spaceBetweenWords := totalSpaces / (len(words) - 1) // Minimum spaces between words
	extraSpaces := totalSpaces % (len(words) - 1)       // Extra spaces to distribute

	// Create a buffer to build the justified line
	var justifiedLine strings.Builder
	for i, word := range words {
		if i > 0 {
			// Add minimum spaces between words
			justifiedLine.WriteString(strings.Repeat(" ", spaceBetweenWords))
			// Add an extra space for the first few words
			if i <= extraSpaces {
				justifiedLine.WriteString(" ")
			}
		}
		// Add the word
		justifiedLine.WriteString(word)
	}

	return justifiedLine.String()
}

func getTerminalSize() (w, h int, err error) {
	var ws Winsize
	_, _, e := syscall.Syscall6(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&ws)),
		0, 0, 0,
	)
	if e != 0 {
		err = e
		return
	}
	w = int(ws.Col)
	h = int(ws.Row)
	return
}

type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}
