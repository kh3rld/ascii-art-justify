package output

import "fmt"

func PrintUsage() {
	fmt.Print("Usage: go run . [OPTION] [STRING] [BANNER]\n\n")
	fmt.Println("Example: go run . --align=right something standard")
}

func OutputUsage() {
	fmt.Print("Usage: go run . [OPTION] [STRING] [BANNER]\n\n")
	fmt.Println("Example: go run . --output=<fileName.txt> something standard")
}
