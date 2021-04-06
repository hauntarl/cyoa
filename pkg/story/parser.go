package story

import (
	"encoding/json"
	"io"
)

// Parse decodes the json file to type: Story
func Parse(r io.Reader) (story Book, err error) {
	d := json.NewDecoder(r)
	err = d.Decode(&story)
	return
}

// Book contains all the chapters that particular story can have
type Book map[string]Chapter

// Chapter stores the data of a particular chapter in the story
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option stores the available choices the user can make
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
