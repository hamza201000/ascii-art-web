package asciart

import (
	"strings"
)
// Split_with_new_line splits the input string by newline characters and returns a slice of lines.
func Split_with_new_line(str string) ([]string, error) {
	word := []string{}
	b := false
	temp_str := ""
	var err error

	str = strings.ReplaceAll(str, "\r\n", "\n") // Replace literal "\n" with actual newline
	for i := 0; i < len(str); i++ {
		if str[i] != '\n' && (str[i] < 32 || str[i] > 126) {
			return nil, err
		}
		if str[i] == '\n' {
			word = append(word, temp_str)
			temp_str = ""
		} else {
			b = true
			temp_str += string(str[i])
		}
	}
	// Add the last word if not empty
	if len(temp_str) > 0 {
		word = append(word, temp_str)
		temp_str = ""
	}
	// Handle case where string ends with newline
	if len(str) > 0 && str[len(str)-1] == '\n' && b {
		word = append(word, temp_str)
	}
	return word, err
}
