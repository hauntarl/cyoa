package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"hauntarl.io/gophercises/cyoa/pkg/story"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA apllication on")
	filename := flag.String("file", "gopher.json", "JSON file for CYOA story.")
	flag.Parse()

	fmt.Printf("Using the story from %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	adventure, err := story.ParseStory(f)
	if err != nil {
		panic(err)
	}

	// creating an handler with default options
	// h := cyoa.NewHandler(story)
	// fmt.Printf("Starting the server on port: %d\n", *port)
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

	// providing a custom template and a custom path parser for that template
	tmpl := template.Must(template.New("").Parse(customTmpl))
	h := story.NewHandler(
		adventure,
		story.WithTemplate(tmpl),
		story.WithPathParser(customPathParser),
	)
	mux := http.NewServeMux()
	mux.Handle("/story/", h)

	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

// custom pathParser
func customPathParser(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

var customTmpl = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      {{if .Options}}
        <ul>
        {{range .Options}}
          <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
        </ul>
      {{else}}
        <h3>The End</h3>
      {{end}}
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`
