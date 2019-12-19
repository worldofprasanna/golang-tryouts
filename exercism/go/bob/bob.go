// Package bob contains the functions which represents bob s character
package bob

import (
	"strings"
	"regexp"
)

// IsUpperCase Letter will check if the string contains atleast one Uppercase character
var IsUpperCase = regexp.MustCompile(`[A-Z]+`).MatchString

// Hey will be used to communicate to bob
func Hey(remark string) string {
	remark = strings.Trim(remark, " \t\n\r")
	switch {
	case remark == "":
		return "Fine. Be that way!"
	case strings.ToUpper(remark) == remark && remark[len(remark) - 1:] == "?" && IsUpperCase(remark):
		return "Calm down, I know what I'm doing!"
	case strings.ToUpper(remark) == remark && IsUpperCase(remark):
		return "Whoa, chill out!"
	case remark[len(remark) - 1:] == "?":
		return "Sure."
	default:
		return "Whatever."
	}
}
