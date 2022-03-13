package parser

import (
	"fmt"
	"regexp"
)

//
// Return the opening, or closing, tag depending on the first arguments truthness
//
func getTag(inside bool, openingTag string, closingTag string) string {
	if inside {
		return closingTag
	}
	return openingTag
}

//
// Parse a line by checking the first character of the string
//
func Parse(str string) string {
	buffer := ""

	// Headers
	if str[0] == '#' {
		n := 1

		for i := 1; i < 6; i++ {
			if str[i] == '#' {
				n++
			}
		}

		buffer += fmt.Sprintf("<h%d>%s</h%[1]d>", n, ParseBody(str[n+1:]))
		return buffer
	}

	// Lists
	if str[0] == '-' {
		list := regexp.MustCompile(`\n`).Split(str, -1)

		buffer += "<ul>"
		for _, l := range list {
			if l != "" {
				buffer += "<li>" + l[2:] + "</li>"
			}
		}
		buffer += "</ul>"
		return buffer
	}

	// Paragraphs
	return "<p>" + ParseBody(str) + "</p>"
}

//
// Parse the line body
//
func ParseBody(str string) string {
	buffer := ""
	insideBold := false
	insideItalic := false
	insideBoldItalic := false

	for i := 0; i < len(str); i++ {
		c := str[i]

		// Styling
		if c == '*' {
			if i+1 < len(str) && str[i+1] == '*' {
				// Bold + italic
				if i+2 < len(str) && str[i+2] == '*' {
					buffer += getTag(insideBoldItalic, "<b><i>", "</i></b>")
					insideBoldItalic = !insideBoldItalic
					i += 2
					continue
				}

				// Bold
				buffer += getTag(insideBold, "<b>", "</b>")
				insideBold = !insideBold
				i++
				continue
			}

			// Italic
			buffer += getTag(insideItalic, "<i>", "</i>")
			insideItalic = !insideItalic
			continue
		}

		buffer += string(c)
	}

	return buffer
}
