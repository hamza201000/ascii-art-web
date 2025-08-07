package asciart

import (
	"bufio"
	"fmt"
	"os"
)

func Fmain(s, Banner string) (string, int) {
	// Check if the correct number of arguments is provided
	// if len(os.Args) != 2 {
	// 	fmt.Println("Usage : go run main.go \"Your text here\"")
	// 	return
	// }
	// Try to open the font files in order: standard, shadow, thinkertoy
	if Banner != "standard" && Banner != "shadow" && Banner != "thinkertoy" {
		return "", 400
	}

	file, err := os.Open("Banner/" + Banner + ".txt")
	if err != nil {
		fmt.Println("Error: failed to open any files")
		return "", 400
	}
	// Ensure the file is closed after finishing
	defer file.Close()

	// Create a scanner to read the font file
	Scanner := bufio.NewScanner(file)

	// Parse the ASCII art table from the font file
	asci_table := ParseAsci(Scanner)
	// Split the input string by new lines
	newstring, err := Split_with_new_line(s)
	if err != nil {
		return "", 400
	}
	// fmt.Println(s)
	// newstring := strings.Split(s, "\n")

	// Print the ASCII art for the input string

	return PrintAsci(newstring, asci_table), 200
}
