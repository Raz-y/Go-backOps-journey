// Package cyoa provides the core logic and structures for creating a Choose Your Own Adventure (CYOA) web application.
// It includes types and functions for loading a story from a JSON file, handling HTTP requests, and rendering story chapters using HTML templates.

package cyoa

import (
	"encoding/json"
	"io"
)

// Story represents a story, which is a map where the keys are chapter names and the values are the chapters themselves.
type Story map[string]Chapter

// Option represents an option inside a Chapter, linking to another chapter in the story.
type Option struct {
	Text    string `json:"text"`
	NextArc string `json:"arc"`
}

// Chapter represents a chapter in a choose your own adventure story.
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

// JSONStory parses a JSON-encoded story from an io.Reader and returns it as a Story.
func JSONStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)

	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}
