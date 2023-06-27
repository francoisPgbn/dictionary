package dictionary

import (
	"time"
)

type Entry struct {
	Definition   string `json:"definition"`
	CreationDate string `json:"creationDate"`
}

func (e Entry) String() string {
	return e.Definition + " " + e.CreationDate
}

type Dictionary struct {
	entries map[string]Entry
}

func New() *Dictionary {
	dictionary := Dictionary { entries: make(map[string]Entry)}
	return &dictionary
}

func (d *Dictionary) Add(word string, definition string) {
	d.entries[word] = Entry{Definition: definition, CreationDate: time.Now().Format("2006-01-02")}
}

func (d *Dictionary) Get(word string) (Entry, error) {
	return d.entries[word], nil
}

func (d *Dictionary) Remove(word string) {
	delete(d.entries, word)
}

func (d *Dictionary) List() ([]string, map[string]Entry) {

	return []string{}, d.entries
}