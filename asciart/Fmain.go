package asciart

import (
	"bufio"
	"net/http"
	"os"
)

func Fmain(s, banner string) (string, int) {
	// Try to open the font files in order: standard, shadow, thinkertoy
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		return "", http.StatusBadRequest
	}

	file, err := os.Open("banner/" + banner + ".txt")
	if err != nil {
		return "", http.StatusInternalServerError
	}
	// Ensure the file is closed after finishing
	defer file.Close()

	// Create a scanner to read the font file
	Scanner := bufio.NewScanner(file)

	// Parse the ASCII art table from the font file
	asci_table := ParseAsci(Scanner)
	if len(asci_table) == 0 {
		return "", http.StatusInternalServerError
	}
	// Split the input string by new lines
	newstring, err := Split_with_new_line(s)
	if err != nil {
		return string(err.Error()), 200
	}

	// Print the ASCII art for the input string

	return PrintAsci(newstring, asci_table), 200
}
