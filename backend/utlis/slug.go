package utils

import (
	"strings"
	"unicode"
)

func Slugify(s string) string {
	/*
		function that takes the title of a blog post and creates a slug
		by replacing spaces with dashes
	*/
	s = strings.ToLower(s)
	var b strings.Builder
	prevDash := false

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			b.WriteRune(r)
			prevDash = false
		} else if prevDash == false {
			b.WriteRune('-')
			prevDash = true
		}
	}
	return strings.Trim(b.String(), "-")
}
