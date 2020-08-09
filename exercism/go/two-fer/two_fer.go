
// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	person := "you"
	if name != "" {
		person = name
	}
	return fmt.Sprintf("One for %s, one for me.", person)
}
