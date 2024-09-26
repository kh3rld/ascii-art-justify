package reading

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"justify/check"
)

// Reading reads the content of the banner file and returns a slice of strings.
// It also validates the integrity of the banner files using checksums.
func Reading(bannerFile string) []string {
	// Check if the file extension is .txt
	if filepath.Ext(bannerFile) != ".txt" {
		fmt.Println("Error: Incorrect file extension for banner file")
		os.Exit(1)
	}

	// Read the file data
	bannerFileData, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Validate the file's checksum
	fileHash := check.ValidFile(bannerFileData)

	// Check for file corruption based on the expected hash values
	expectedHashes := map[string]string{
		"standard.txt":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
		"shadow.txt":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
		"thinkertoy.txt": "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3",
	}

	if expectedHash, exists := expectedHashes[bannerFile]; exists {
		if fileHash != expectedHash {
			fmt.Printf("Error: The banner file \"%s\" is corrupted\n", bannerFile)
			os.Exit(1)
		}
	}

	// Split the banner file data into lines
	var splitBannerFileData []string
	if bannerFile == "thinkertoy.txt" {
		splitBannerFileData = strings.Split(string(bannerFileData), "\r\n")
	} else {
		splitBannerFileData = strings.Split(string(bannerFileData), "\n")
	}

	return splitBannerFileData
}
