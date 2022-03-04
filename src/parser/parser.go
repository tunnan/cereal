package parser

import (
	"regexp"
	"strings"
)

//
// Token struct
//
type token struct {
  Tag string
  Body string
}

var CRLF = regexp.MustCompile("\r?\n")
var CRLFCRLF = regexp.MustCompile("\r?\n\r?\n")

//
// Return a token with the proper tag and body
// based on the matching prefix
//
func determineToken(line string) (t *token) {
  m := map[string]string{
    "# ": "h1",
    "!": "img",
    "[": "a",
    "- ": "ul",
    "* ": "ol",
  }

  for prefix, tag := range m {
    if strings.HasPrefix(line, prefix) {
      t = &token{tag, line[len(prefix):]}
    }
  }

  if t == nil {
    t = &token{"p", line}
  }

  return
}

//
// Tokenize every line into an array of tokens
//
func tokenize(lines []string) (tokens []token) {
  for _, line := range lines {
    tokens = append(tokens, *determineToken(line))
  }

  return
}

//
// Parse the body into HTML
//
func Parse(body string) (html string){
  lines := CRLFCRLF.Split(body, -1)
  tokens := tokenize(lines) 

  for _, token := range tokens {
    switch token.Tag {

    // Images
    case "img":
      replacer := strings.NewReplacer(
        "[", "<img alt=\"",
        "](", "\" src=\"",
        ")", "\">")
      html += replacer.Replace(token.Body)

    // Links
    case "a":
      innerText := token.Body[:strings.Index(token.Body, "]")]
      href := token.Body[strings.Index(token.Body, "(")+1:strings.Index(token.Body, ")")]
      html += "<a href=\""+href+"\">"+innerText+"</a>"

    // Unordered lists
    case "ul":
      html += "<ul>"
      for _, li := range CRLF.Split("- " + token.Body, -1) {
        v := strings.TrimPrefix(li, "- ")
        if v != "" {
          html += "<li>"+v+"</li>"
        }
      }
      html += "</ul>"

    // Ordered lists
    case "ol":
      html += "<ol>"
      for _, li := range CRLF.Split("* " + token.Body, -1) {
        v := strings.TrimPrefix(li, "* ")
        if v != "" {
          html += "<li>"+v+"</li>"
        }
      }
      html += "</ol>"

    // Default to <p>
    default:
      html += "<"+token.Tag+">"+token.Body+"</"+token.Tag+">"
    }
  }

  return
}
