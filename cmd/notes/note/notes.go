package note

import (
	"io/ioutil"

	"go-mod.ewintr.nl/shitty-ssg/pkg/adoc"
)

type Notes []*Note

func (n *Notes) AddFileNote(fPath string) error {
	content, err := ioutil.ReadFile(fPath)
	if err != nil {
		return err
	}
	note := NewNote(adoc.New(string(content)))
	if note.Kind != KIND_INVALID {
		*n = append(*n, note)
	}

	return nil
}

func (n *Notes) FilterByTerm(term string) Notes {
	found := Notes{}
	for _, note := range *n {
		if note.Contains(term) {
			found = append(found, note)
		}
	}

	return found
}
