package story

import (
	"encoding/json"
	"io"
)

// ParseStory decodes the json file to type: Story
func ParseStory(r io.Reader) (story Story, err error) {
	d := json.NewDecoder(r)
	err = d.Decode(&story)
	return
}

// Story contains all the chapters that particular story can have
type Story map[string]Chapter

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
