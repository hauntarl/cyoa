package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"hauntarl.io/gophercises/cyoa/pkg/story"
)

var (
	fname *string
	first *string
)

func init() {
	fname = flag.String("file", "gopher.json", "JSON file for CYOA story.")
	first = flag.String("start", "intro", "The first chapter of the story")
	flag.Parse()
	log.Printf("using the story from file %s.\n", *fname)
}

func main() {
	f, err := os.Open(*fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	book, err := story.Parse(f)
	if err != nil {
		panic(err)
	}

	err = begin(book, *first)
	if err != nil {
		fmt.Println(err)
	}
}

func begin(book story.Book, choice string) error {
	var count int
	for choice != "" {
		count++
		chapter, err := fetchNext(&book, choice)
		if err != nil {
			return err
		}
		read(chapter, count)
		choice = makeAChoice(chapter)
	}
	return nil
}

func fetchNext(book *story.Book, id string) (*story.Chapter, error) {
	if chapter, ok := (*book)[id]; ok {
		return &chapter, nil
	}
	return nil, fmt.Errorf("chapter '%s' not found", id)
}

func read(chapter *story.Chapter, num int) {
	fmt.Printf("\n\nChapter %d: %s\n", num, chapter.Title)
	for _, paragraph := range chapter.Paragraphs {
		fmt.Printf("\n%s", paragraph)
	}

	fmt.Println("\n\nHow'd you like to respond?")
	for i, option := range chapter.Options {
		fmt.Printf("%d. %s\n", i+1, option.Text)
	}
}

func makeAChoice(chapter *story.Chapter) string {
	var res int
	fmt.Scanf("%d\n", &res)
	if res < 1 || res > len(chapter.Options) {
		fmt.Println("\nThanks for Reading!")
		return ""
	}
	return chapter.Options[res-1].Chapter
}
