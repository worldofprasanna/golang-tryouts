// Package acronym contains functions to operate on acronyms
package acronym

import (
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	replacer := strings.NewReplacer("_", " ", ",", " ", "-", " ")
	trimmedStr := replacer.Replace(s)
	trimmedStrArr := strings.Split(trimmedStr, " ")
	result := ""
	for _, word := range trimmedStrArr {
		result += firstCharOfWord(word)
	}
	return strings.ToUpper(result)
}

func firstCharOfWord(word string) string {
	word = strings.Trim(word, " ")
	if len(word) == 0 {
		return ""
	} 
	return string(word[0])
}
