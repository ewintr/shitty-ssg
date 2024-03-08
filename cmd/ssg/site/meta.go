package site

import (
	"code.ewintr.nl/shitty-ssg/pkg/adoc"
)

const (
	KIND_NOTE    = Kind("note")
	KIND_STORY   = Kind("story")
	KIND_ARTICLE = Kind("article")
	KIND_INVALID = Kind("")
)

type Kind string

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
