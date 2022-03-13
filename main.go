package main

import (
	"os"
  "regexp"

	"github.com/tunnan/cereal/src/parser"
	"github.com/tunnan/cereal/src/util"
)

// TODO
// - Add some way to render navigation or something
// - Only parse files if they have been changed since last time
// - Handle all potential errors

func main() {
  entries, _ := os.ReadDir("./app/pages")

  for _, name := range util.GetMarkdownFiles(entries) {
    contents, _ := os.ReadFile("./app/pages/" + name + ".md")
	  contents = regexp.MustCompile(`\r\n`).ReplaceAll(contents, []byte("\n"))
  	lines := regexp.MustCompile(`\n\n`).Split(string(contents), -1)
    
    html := ""
    for _, line := range lines {
      //fmt.Printf("%d: %s\n", i, parser.Parse(line))
      html += parser.Parse(line)
    }

    os.WriteFile("./dist/" + name + ".html", []byte(util.WrapInTemplate(html)), 0600)
  }
}
