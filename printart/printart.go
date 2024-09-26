package printart

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

func PrintArt(bannerFileSlice []string, inputString string, alignFlag string) {
	// Determine alignment
	align := "left"
	if alignFlag != "" {
		align = strings.ToLower(alignFlag[8:])
		if !isValidAlignment(align) {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
			return
		}
	}

	if align == "justify" {
		inputString = strings.Join(strings.Fields(inputString), " ")
	}

	// fmt.Println(align)
	// fmt.Println(alignFlag)

	if inputString == "\\n" {
		fmt.Println()
		return
	}
	// Check for unprintable sequences
	if checkUnprintableSequences(inputString) {
		return
	}

	// Replace escape sequences
	processedInput := processEscapeSequences(inputString)

	// Split input into lines
	splitArguments := strings.Split(processedInput, "\n")
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
		width, _, err := getTerminalSize() // Get terminal size
		if err != nil {
			fmt.Println("Error getting terminal size:", err)
			return
		}

		art := generateArt(bannerFileSlice, text, asciiHeight)

		result := art
		alignedLines := []string{}
		switch align {
		case "left":
			fmt.Println(result)
		case "right":
			artSlice := strings.Split(result, "\n")
			artLen := len(artSlice[0])

			spacesRem := width - artLen

			if spacesRem < 0 {
				fmt.Println(result)
				return
			}
			for _, line := range artSlice {
				alignedLines = append(alignedLines, strings.Repeat(" ", spacesRem)+line)
			}
			fmt.Println(strings.Join(alignedLines, "\n"))

		case "center":
			artSlice := strings.Split(result, "\n")
			artLen := len(artSlice[0])

			spacesRem := width - artLen

			if spacesRem < 0 {
				fmt.Println(result)
				return
			}
			for _, line := range artSlice {
				alignedLines = append(alignedLines, strings.Repeat(" ", spacesRem/2)+line)
			}
			fmt.Println(strings.Join(alignedLines, "\n"))
		case "justify":

			inputSlice := strings.Fields(inputString)
			inputString = strings.Join(inputSlice, " ")
			spacePositions, asciiArtLines := spacePos(inputString, bannerFileSlice)
			artSlice := strings.Split(result, "\n")
			artLen := len(artSlice[0])

			if artLen < width && len(spacePositions) > 0 {
				extraSpaces := width - artLen
				spaceToAdd := extraSpaces / len(spacePositions)
				rem := extraSpaces % len(spacePositions)

				for _, line := range asciiArtLines {
					newLine := []rune(line)
					offset := 0

					for i, pos := range spacePositions {
						additionalSpace := spaceToAdd

						if i < rem {
							additionalSpace++
						}
						adjustedPosition := pos + offset
						newLine = append(newLine[:adjustedPosition], append([]rune(strings.Repeat(" ", additionalSpace)), newLine[adjustedPosition:]...)...)
						offset += additionalSpace
					}
					alignedLines = append(alignedLines, string(newLine))
				}
			} else {
				alignedLines = asciiArtLines[:]
			}
			fmt.Println(strings.Join(alignedLines, "\n"))
		}
	}
}

// Checks if the alignment string is valid
func isValidAlignment(align string) bool {
	validAlignments := map[string]bool{"left": true, "right": true, "justify": true, "center": true}
	return validAlignments[align]
}

// Processes special input cases.
func handleSpecialInputs(input string) bool {
	switch input {
	case "\\n":
		fmt.Println()
		return true
	case "":
		return true
	case "\\t":
		fmt.Println("    ")
		return true
	}
	return false
}

// Checks for unprintable character sequences in input.
func checkUnprintableSequences(inputString string) bool {
	unprintableSequences := []string{"\\a", "\\b", "\\v", "\\f", "\\r"}
	for _, unprintable := range unprintableSequences {
		if strings.Contains(inputString, unprintable) {
			fmt.Println("Input string contains an unprintable sequence")
			return true
		}
	}
	return false
}

// Replaces escape sequences with their actual representations.
func processEscapeSequences(inputString string) string {
	tabCharText := strings.Replace(inputString, "\\t", "    ", -1)
	return strings.ReplaceAll(tabCharText, "\\n", "\n")
}

func spacePos(input string, reading []string) ([]int, [8]string) {
	spacePosition := []int{}
	totalTextWidth := 0
	asciiArtLines := [8]string{}

	for _, cha := range input {
		startingIndex := int(cha-32)*9 + 1
		art := reading[startingIndex : startingIndex+8]

		for i := 0; i < 8; i++ {
			if i < len(art) {
				asciiArtLines[i] += art[i]
			} else {
				asciiArtLines[i] += strings.Repeat(" ", len(art[0]))
			}
		}
		if cha == ' ' {
			spacePosition = append(spacePosition, totalTextWidth)
		}

		totalTextWidth += len(art[0])
	}

	return spacePosition, asciiArtLines
}

// generateArt creates ASCII art for the given text.
func generateArt(bannerFileSlice []string, text string, asciiHeight int) string {
	var art strings.Builder
	for j := 0; j < asciiHeight; j++ {
		for _, char := range text {
			startingIndex := int(char-32)*9 + 1
			art.WriteString(bannerFileSlice[startingIndex+j])
		}
		art.WriteString("\n")
	}
	return art.String()
}

// Winsize structure for terminal window size.
type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

// Retrieves the terminal window size.
var syscallGetWinsize = func(ws *Winsize) (uintptr, uintptr, syscall.Errno) {
	return syscall.Syscall(syscall.SYS_IOCTL,

		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)
}

// Returns the width and height of the terminal
func getTerminalSize() (int, int, error) {
	ws := &Winsize{}
	_, _, err := syscallGetWinsize(ws)
	if err != 0 {
		return 0, 0, err
	}
	return int(ws.Col), int(ws.Row), nil
}
