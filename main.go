package main

import (
	"os"
  "regexp"
  "strings"

	"github.com/tunnan/cereal/src/parser"
)

// TODO
// - Add some way to render navigation or something
// - Only parse files if they have been changed since last time
// - Handle all potential errors
// - Remove the temporary favicon fix in the template

//
// Filter out all the files not ending with .md
//
func getMarkdownFiles(entries []os.DirEntry) (files []string) {
  for _, entry := range entries {
    if (!strings.HasSuffix(entry.Name(), ".md")) {
      continue
    }
    files = append(files, entry.Name()[:len(entry.Name()) - 3])
  }

  return
}

//
// Main
//
func main() {
  entries, _ := os.ReadDir("./app/pages")

  // Go through all the md files in the "app/page" directory
  for _, name := range getMarkdownFiles(entries) {
    contents, _ := os.ReadFile("./app/pages/" + name + ".md")
	  contents = regexp.MustCompile(`\r\n`).ReplaceAll(contents, []byte("\n"))
  	lines := regexp.MustCompile(`\n\n`).Split(string(contents), -1)

    // Get the body HTML of the page
    page := ""
    for _, line := range lines {
      page += parser.Parse(line)
    }

    // Wrap the body in the template HTML
    html := `<!DOCTYPE html>
      <html lang="en">
      <head>
        <meta charset="utf-8">
        <!--<link rel="icon" type="image/png" href="./favicon.png">-->
        <link rel="shortcut icon" href="data:image/x-icon;," type="image/x-icon"> 
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <link href="https://fonts.googleapis.com/css2?family=Poppins:wght@400;500;600&display=swap" rel="stylesheet"> 
        <title>Cereal</title>
        <link rel="stylesheet" href="./assets/style.css">
      </head>
      <body>
        <div id="app">` + page + `</div>
      </body>
      </html>`

    // Write to file
    os.WriteFile("./dist/" + name + ".html", []byte(html), 0600)
  }
}
