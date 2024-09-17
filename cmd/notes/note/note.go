package note

import (
	"strings"

	"go-mod.ewintr.nl/shitty-ssg/pkg/adoc"
)

type Kind string

type Tag string

const (
	KIND_NOTE         = Kind("note")
	KIND_PRIVATE_NOTE = Kind("private_note")
	KIND_WORK_NOTE    = Kind("work_note")
	KIND_INVALID      = Kind("")
)

func mapKind(akind adoc.Kind) Kind {
	nkind, ok := map[adoc.Kind]Kind{
		adoc.KIND_NOTE:         KIND_NOTE,
		adoc.KIND_PRIVATE_NOTE: KIND_PRIVATE_NOTE,
		adoc.KIND_WORK_NOTE:    KIND_WORK_NOTE,
	}[akind]
	if !ok {
		return KIND_INVALID
	}

	return nkind
}

type Note struct {
	doc     *adoc.ADoc
	Title   string
	Kind    Kind
	Tags    []Tag
	Content string
}

func NewNote(doc *adoc.ADoc) *Note {
	var paragraphs []string
	for _, be := range doc.Content {
		paragraphs = append(paragraphs, be.Text())
	}
	content := strings.Join(paragraphs, "\n\n")

	var tags []Tag
	for _, t := range doc.Tags {
		tags = append(tags, Tag(t))
	}

	return &Note{
		doc:     doc,
		Kind:    mapKind(doc.Kind),
		Title:   doc.Title,
		Tags:    tags,
		Content: content,
	}
}

func (n *Note) Contains(term string) bool {
	for _, t := range n.Tags {
		if strings.ToLower(string(t)) == strings.ToLower(term) {
			return true
		}
	}

	for _, w := range strings.Split(n.Content, " ") {
		if strings.ToLower(w) == strings.ToLower(term) {
			return true
		}
	}

	return false
}
