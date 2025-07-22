package asciart

// PrintAsci prints the ASCII art for each line in str using asci_table.
func PrintAsci(str []string, asci_table [][]string) string {
	printit := ""
	for k := 0; k < len(str); k++ {
		for i := 0; i < 8; i++ { // Each character is 8 lines tall
			for j := 0; j < len(str[k]); j++ {
				m := (str[k][j] - 32) // Map character to table index

				printit += (asci_table[m][i])

			}
			printit += "\n" // Move to the next line after printing one row of ASCII art for all characters
			if len(str[k]) == 0 {
				break // Skip empty lines
			}
		}
	}
	return printit
}
