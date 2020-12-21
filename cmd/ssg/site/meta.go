package site

import (
	"git.sr.ht/~ewintr/shitty-ssg/pkg/adoc"
)

const (
	KIND_NOTE    = Kind("note")
	KIND_STORY   = Kind("story")
	KIND_ARTICLE = Kind("article")
	KIND_INVALID = Kind("")
)

type Kind string

func NewKind(kind adoc.Kind) Kind {
	switch kind {
	case adoc.KIND_NOTE:
		return KIND_NOTE
	case adoc.KIND_VKV:
		return KIND_STORY
	case adoc.KIND_ESSAY:
		fallthrough
	case adoc.KIND_TUTORIAL:
		return KIND_ARTICLE
	default:
		return KIND_INVALID
	}

}

const (
	LANGUAGE_EN      = Language("en")
	LANGUAGE_NL      = Language("nl")
	LANGUAGE_INVALID = Language("")
)

type Language string

func NewLanguage(ln adoc.Language) Language {
	switch ln {
	case adoc.LANGUAGE_NL:
		return LANGUAGE_NL
	case adoc.LANGUAGE_EN:
		fallthrough
	default:
		return LANGUAGE_EN
	}
}

type Tag string

const (
	DATE_SHORT DateFormat = iota
	DATE_LONG
	DATE_FORMAL //Sat, 12 Sep 2020 12:32:00 +0200
)

type DateFormat int
