package site

const (
	KIND_NOTE    = Kind("note")
	KIND_STORY   = Kind("story")
	KIND_ARTICLE = Kind("article")
	KIND_INVALID = Kind("")
)

type Kind string

var pluralKind = map[Kind]string{
	KIND_NOTE:    "notes",
	KIND_STORY:   "stories",
	KIND_ARTICLE: "articles",
}

func NewKind(kind string) Kind {
	switch kind {
	case "note":
		return KIND_NOTE
	case "story":
		return KIND_STORY
	case "article":
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

func NewLanguage(text string) Language {
	switch text {
	case "nl":
		return LANGUAGE_NL
	case "en":
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
