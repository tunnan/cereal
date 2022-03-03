package util

import (
	"os"
	"strings"
)

//
// Filter out all the files not ending with .md
//
func GetMarkdownFiles(entries []os.DirEntry) (files []string) {
  for _, entry := range entries {
    if (!strings.HasSuffix(entry.Name(), ".md")) {
      continue
    }

    files = append(files, entry.Name()[:len(entry.Name()) - 3])
  }

  return
}

//
// Wrap the page contents in a template
//
func WrapInTemplate(html string) string {
  template, _ := os.ReadFile("./app/templates/default.html")
  return strings.Replace(string(template), "{{main}}", html, 1)
}
