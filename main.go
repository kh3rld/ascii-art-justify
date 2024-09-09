package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	check "justify/checksum"
	print "justify/printAscii"
	output "justify/readWrite"
	usage "justify/utils"
)

var banners = map[string]string{
	"standard":   "standard.txt",
	"thinkertoy": "thinkertoy.txt",
	"shadow":     "shadow.txt",
}

func main() {
	checksum := flag.Bool("checksum", false, "Check integrity of specified file")
	flname := flag.String("output", "", "Usage: go run . [OPTION] [STRING] [BANNER]")
	align := flag.String("align", "left", "Alignment type: center, left, right, justify")
	flag.Parse()

	// Validate the align option
	if *align != "center" && *align != "left" && *align != "right" && *align != "justify" {
		usage.PrintUsage()
		return
	}

	var presentF, correctF bool
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "output" {
			presentF = true
			result := strings.Replace(os.Args[1], *flname, "", 1)
			if !(result == "--output=") {
				correctF = true
			}
		}
	})

	if presentF && correctF {
		usage.PrintUsage()
		return
	}

	args := flag.Args()

	if len(args) == 0 {
		usage.PrintUsage()
		return
	}

	input := args[0]
	var banner string

	if len(args) > 1 {
		banner = args[1]
	} else {
		banner = "standard"
	}

	filename, ok := banners[banner]
	if !ok {
		fmt.Println("Invalid banner specified.")
		return
	}
	if *checksum {
		err := check.ValidateFileChecksum(filename)
		if err != nil {
			log.Fatalf("Error checking integrity: %v", err)
		}
		fmt.Printf("Integrity check passed for file: %s\n", filename)
		return
	}

	err := check.ValidateFileChecksum(filename)
	if err != nil {
		log.Printf("Error downloading or validating file: %v", err)
		return
	}

	asciiArtGrid, err := output.ReadAscii(filename)
	if err != nil {
		log.Fatalf("Error reading ASCII map: %v", err)
	}

	data, err := print.WriteArt(input, asciiArtGrid)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	if *flname != "" {
		filename := strings.TrimPrefix(*flname, "--output=")
		if filename == "" {
			fmt.Println("Error: --output flag must be followed by a filename")
			usage.PrintUsage()
			os.Exit(1)
		}
		err = output.WriteAscii(data, filename)
		if err != nil {
			log.Printf("error: %v", err)
		}
	} else {
		err = print.PrintArt(input, asciiArtGrid, *align)
		if err != nil {
			log.Printf("Error: %v", err)
		}
	}
}
