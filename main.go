package main

import (
	"os"
	"github.com/tunnan/cereal/src/parser"
	"github.com/tunnan/cereal/src/util"
)

// TODO
// - Handle parsing for lists, etc..
// - Only parse files if they have been changed since last time
// - Handle all potential errors

func main() {
  entries, _ := os.ReadDir("./app/pages")

  for _, name := range util.GetMarkdownFiles(entries) {
    contents, _ := os.ReadFile("./app/pages/" + name + ".md")
    html := parser.Parse(string(contents))

    os.WriteFile("./dist/" + name + ".html", []byte(util.WrapInTemplate(html)), 0600)
  }
}
