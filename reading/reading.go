package reading

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"ascii-art-justify/check"
)

// Reading reads the content of the banner file and returns a slice of string.
// It also implements a checksum validator that checks for the integrity of the banner files.
func Reading(bannerFile string) []string {
	if filepath.Ext(bannerFile) != ".txt" {
		fmt.Println("Incorrect file extension associated with banner file")
		os.Exit(1)
	}

	bannerFileData, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("Error reading file", err)
		os.Exit(1)
	}

	fileHash := check.ValidFile(bannerFileData)

	switch bannerFile {
	case "standard.txt":
		if fileHash != "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf" {
			fmt.Println("error: the banner file \"standard.txt\" is corrupted")
			os.Exit(1)
		}
	case "shadow.txt":
		if fileHash != "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73" {
			fmt.Println("error: the banner file \"shadow.txt\" is corrupted")
			os.Exit(1)
		}
	case "thinkertoy.txt":
		if fileHash != "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3" {
			fmt.Println("error: the banner file \"thinkertoy.txt\" is corrupted")
			os.Exit(1)
		}

	}

	var splitBannerFileData []string

	if bannerFile == "thinkertoy.txt" {
		splitBannerFileData = strings.Split(string(bannerFileData), "\r\n")
	} else {
		splitBannerFileData = strings.Split(string(bannerFileData), "\n")
	}

	return splitBannerFileData
}
