package asciart

import "bufio"

// ParseAsci parses the font file and returns a 2D slice of ASCII art characters.
func ParseAsci(Scanner *bufio.Scanner) [][]string {
	asci := [][]string{}
	asci_line := []string{}
	b := false // Flag to skip leading empty lines

	for Scanner.Scan() {
		// Skip leading empty lines before the first character
		if Scanner.Text() == "" && !b {
			continue
		}
		// If an empty line is found, it marks the end of a character
		if Scanner.Text() == "" {
			asci = append(asci, asci_line)
			asci_line = []string{}
		} else {
			asci_line = append(asci_line, Scanner.Text())
			b = true
		}
	}
	if len(asci_line) > 0 {
		asci = append(asci, asci_line) // Append the last character if not empty
	}
	return asci
}
