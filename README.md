# Choose your own adventure

Choose Your Own Adventure is (was?) a series of books intended for children where as you read you would occasionally be given options about how you want to proceed. For instance, you might read about a boy walking in a cave when he stumbles across a dark passage or a ladder leading to an upper level and the reader will be presented with two options like:

- Turn to page 44 to go up the ladder.
- Turn to page 87 to venture down the dark passage.

The goal is to recreate this experience via a web application where each page will be a portion of the story, and at the end of every page the user will be given a series of options to choose from (or be told that they have reached the end of that particular story arc).

Implementation of Choose your own adventure from gophercises, including the bonus section.

**[Gophercises](https://courses.calhoun.io/courses/cor_gophercises)**  by Jon Calhoun

**Run Commands:**

- go run cmd\web\main.go (for running the web application)
- command-line flags: --port, --file, --help/-h
* go run cmd\terminal\main.go (for running the command-line application)
* command-line flags: --start, --file, --help/-h

**Features:**

- Parsing JSON file: **[Json To Go](https://mholt.github.io/json-to-go/)**
- Using html templates to display the story
- Using http.Handler interface to serve the incoming requests
- Allowing user to pass custom templates and/or path parser
* Command-line application allows user to pass a custom starting point of the story

**Packages explored:**

- "fmt"
- "flag" - command line inputs
- "html/template" - to render templates
- "log" - to log events
- "os" - to open the json file
- "net/http" - to create server
