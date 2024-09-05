package printart

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

func PrintArt(bannerFileSlice []string, inputString string, alignFlag string) {
	align := ""
	if alignFlag != "" {
		align = strings.ToLower(alignFlag[8:])
	} else {
		align = "left"
		// fmt.Println("align should have a value")
	}

	// fmt.Println(align)

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

		width, _, err := getTerminalSize()
		if err != nil {
			fmt.Println("Error getting terminal size:", err)
			return
		}
		terminalSize := width
		var art strings.Builder

		// fmt.Println(terminalSize)
		for j := 0; j < asciiHeight; j++ {
			for _, char := range text {

				startingIndex := int(char-32)*9 + 1
				art.WriteString(bannerFileSlice[startingIndex+j])
			}
			art.WriteString("\n")
		}
		result := art.String()
		alignedLines := []string{}
		switch align {
		case "left":
			fmt.Println(result)
		case "right":
			artSlice := strings.Split(result, "\n")
			artLen := len(artSlice[0])
			// fmt.Println(artLen)
			spacesRem := terminalSize - artLen
			// fmt.Println(artLen)

			if spacesRem < 0 {
				fmt.Println(result)
				return
			}
			for _, line := range artSlice {
				alignedLines = append(alignedLines, strings.Repeat(" ", spacesRem)+line)
			}
			fmt.Println(strings.Join(alignedLines, "\n"))
			// fmt.Printf("%q", repeated)

			// case "center":
			// 	for _, char := range text {

			// 		startingIndex := int(char-32)*9 + 1
			// 		fmt.Printf(bannerFileSlice[startingIndex+j])
			// 	}
			// 	fmt.Println()
			// case "justify":
			// 	for _, char := range text {

			// 		startingIndex := int(char-32)*9 + 1
			// 		fmt.Printf(bannerFileSlice[startingIndex+j])
			// 	}
			// 	fmt.Println()
		}
	}
}

type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

var syscallGetWinsize = func(ws *Winsize) (uintptr, uintptr, syscall.Errno) {
	return syscall.Syscall(syscall.SYS_IOCTL,

		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)
}

func getTerminalSize() (int, int, error) {
	ws := &Winsize{}
	_, _, err := syscallGetWinsize(ws)
	if err != 0 {
		return 0, 0, err
	}
	return int(ws.Col), int(ws.Row), nil
}
