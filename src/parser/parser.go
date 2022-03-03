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

//
// Split the body contents at every empty line (\n\n)
//
func splitBody(body string) (result []string) {
  newline := regexp.MustCompile("\r?\n\r?\n")
  newlineEnd := regexp.MustCompile("\r?\n$")
  lines := newline.Split(body, -1)

  for _, line := range lines {
    replaced := newlineEnd.ReplaceAll([]byte(line), []byte(""))
    result = append(result, string(replaced))
  }

  return
}

//
// Return a token with the proper tag and body
// based on the matching prefix
//
func determineToken(line string) (t *token) {
  m := map[string]string{
    "# ": "h1",
    "!": "img",
    "[": "a",
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
  tokens := tokenize(splitBody(body))

  for _, token := range tokens {
    switch token.Tag {
    case "img":
      replacer := strings.NewReplacer(
        "[", "<img alt=\"",
        "](", "\" src=\"",
        ")", "\">")
      html += replacer.Replace(token.Body)
    case "a":
      innerText := token.Body[:strings.Index(token.Body, "]")]
      href := token.Body[strings.Index(token.Body, "(")+1:strings.Index(token.Body, ")")]
      html += "<a href=\""+href+"\">"+innerText+"</a>"
    default:
      html += "<"+token.Tag+">"+token.Body+"</"+token.Tag+">"
    }
  }

  return
}
