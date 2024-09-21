package output

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

				// Combine the ASCII art lines into a block of text
				// asciiBlock := strings.Join(artLines, "\n")

				// Align the entire block of ASCII art
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
			// case "justify":
			// alignedLines = append(alignedLines, justifyAlign(line, width))
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

func PrintArtJustify(str string, asciiArtGrid [][]string, align string) error {
	fmt.Println(str)
	terminalWidth, _, err := getTerminalSize()
	fmt.Println(terminalWidth)
	if err != nil {
		return err
	}

	s := strings.ReplaceAll(str, "\\n", "\n")
	s = strings.ReplaceAll(s, "\\r", "\r")
	s = strings.ReplaceAll(s, "\\f", "\f")
	s = strings.ReplaceAll(s, "\\v", "\v")
	s = strings.ReplaceAll(s, "\\t", "\t")
	s = strings.ReplaceAll(s, "\\b", "\b")
	s = strings.ReplaceAll(s, "\\a", "\a")
	lines := strings.Split(s, "\n")
	var justifiedLine string
	for _, line := range lines {

		totalChars := 0
		var r []rune

		for _, x := range line {
			r = append(r, x)
		}

		fmt.Println(string(r))
		for i := range r {
			if r[i] == ' ' {
				r[i] = 'x'
			} else {
				totalChars++
			}
		}
		fmt.Println(string(r))
		words := strings.Fields(string(r))
		if len(words) == 0 {
			fmt.Println()
			continue
		}

		if len(words) == 1 {
			justifiedLine = words[0]
			c := 0
			for _, x := range justifiedLine {
				if x == 'x' {
					c++
				}
			}
			numSpaces := c
			fmt.Println(numSpaces)
			fmt.Println(totalChars)
			fmt.Println(terminalWidth)
			totalSp := terminalWidth - len(asciiArtGrid)
			fmt.Println(totalSp)
			spPad := totalSp / numSpaces
			fmt.Println(spPad)

			for i := range justifiedLine {
				// if x == 'x' {
				// d := strings.ReplaceAll(justifiedLine, "x", " ")
				if i < len(justifiedLine)-1 {
					justifiedLine = strings.ReplaceAll(justifiedLine, "x", strings.Repeat(" ", spPad))
				}

				// }
			}
		} else {

			fmt.Println(totalChars)
			numSpaces := len(words) - 1
			fmt.Println(numSpaces)
			totalSpaces := terminalWidth - totalChars
			spacesBetweenWords := totalSpaces / numSpaces
			// extraSpaces := totalSpaces % numSpaces
			for i, word := range words {
				justifiedLine += word
				if i < len(words)-1 {
					justifiedLine += strings.Repeat("x", spacesBetweenWords)
				}
			}
		}
	}
	fmt.Println(justifiedLine)
	art := asciiArtGrid
	r := ""
	for _, x := range art {
		// fmt.Println(x)
		for _, v := range x {
			for c := range v {
				r = strings.ReplaceAll(string(v[c]), " ", "$")
			}
		}
	}
	fmt.Println(r)
	printWord(justifiedLine, asciiArtGrid)
	return nil
}

func printWord(word string, asciiArtGrid [][]string) {
	for i := 1; i <= 8; i++ {
		for _, char := range word {
			index := int(char - 32)
			if index >= 0 && index < len(asciiArtGrid) {
				fmt.Print(asciiArtGrid[index][i])
				// fmt.Println(len(asciiArtGrid))
			}
		}
		fmt.Println()
	}
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
