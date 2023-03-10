package adoc

import (
	"strings"
	"time"
)

const (
	KIND_NOTE         = Kind("note")
	KIND_PRIVATE_NOTE = Kind("private_note")
	KIND_WORK_NOTE    = Kind("work_note")
	KIND_VKV          = Kind("vkv")
	KIND_STORY        = Kind("story")
	KIND_SNIPPET      = Kind("snippet")
	KIND_ESSAY        = Kind("essay")
	KIND_ARTICLE      = Kind("article")
	KIND_TUTORIAL     = Kind("tutorial")
	KIND_UNKNOWN      = Kind("unknown")
)

type Kind string

func NewKind(text string) Kind {
	switch text {
	case "verhaal":
		text = "story"
	case "los":
		text = "snippet"
	}

	for _, k := range []string{
		"note", "private_note", "work_note", "vkv",
		"story", "snippet",
		"essay", "tutorial", "work", "article",
	} {
		if k == text {
			return Kind(k)
		}
	}

	return KIND_UNKNOWN
}

const (
	LANGUAGE_EN      = Language("en")
	LANGUAGE_NL      = Language("nl")
	LANGUAGE_UNKNOWN = Language("unknown")
)

type Language string

func NewLanguage(ln string) Language {
	switch strings.ToLower(ln) {
	case "nl":
		return LANGUAGE_NL
	case "en":
		return LANGUAGE_EN
	default:
		return LANGUAGE_UNKNOWN
	}
}

type Tag string

type ADoc struct {
	Title    string
	Author   string
	Kind     Kind
	Language Language
	Public   bool
	Path     string
	Date     time.Time
	Tags     []Tag
	Content  []BlockElement
}
