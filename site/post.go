package site

import (
	"errors"
	"fmt"
	"html"
	"path"
	"strconv"
	"strings"
	"time"

	"git.sr.ht/~ewintr/go-kit/slugify"
)

var (
	ErrInvalidPost = errors.New("invalid post")
)

type Post struct {
	Title    string
	Author   string
	Kind     Kind
	Language Language
	Path     string
	Date     time.Time
	Tags     []Tag
	Content  []BlockElement
}

func (p Post) Slug() string {
	return slugify.Slugify(p.Title)
}

func (p *Post) Year() string {
	return strconv.Itoa(p.Date.Year())
}

func (p *Post) Link() string {
	return fmt.Sprintf("%s/", path.Join("/", pluralKind[p.Kind], p.Year(), p.Slug()))
}

func (p *Post) FullLink() string {
	return fmt.Sprintf("https://erikwinter.nl/%s", p.Link())
}

func (p *Post) HTMLSummary() *HTMLSummary {
	summary := ""
	if len(p.Content) > 0 {
		summary = fmt.Sprintf("<p>%s...</p>", truncateOnSpace(p.Content[0].Text(), 150))
	}

	return &HTMLSummary{
		Link:      p.Link(),
		Title:     p.Title,
		Language:  p.Language,
		DateLong:  p.FormattedDate(DATE_LONG),
		DateShort: p.FormattedDate(DATE_SHORT),
		Summary:   summary,
	}
}

func (p *Post) HTMLPost() *HTMLPost {
	var content string
	for _, be := range p.Content {
		content += fmt.Sprintf("%s\n", be.BlockHTML())
	}

	return &HTMLPost{
		Slug:      p.Slug(),
		Title:     html.EscapeString(p.Title),
		DateLong:  p.FormattedDate(DATE_LONG),
		DateShort: p.FormattedDate(DATE_SHORT),
		Content:   content,
	}
}

func (p *Post) XMLPost() *XMLPost {
	var content string
	for _, be := range p.Content {
		content += fmt.Sprintf("%s\n", html.EscapeString(be.BlockHTML()))
	}

	return &XMLPost{
		Link:       p.FullLink(),
		Title:      html.EscapeString(p.Title),
		DateFormal: p.FormattedDate(DATE_FORMAL),
		Content:    content,
	}
}

func (p Post) FormattedDate(format DateFormat) string {
	switch {
	case format == DATE_LONG && p.Language == LANGUAGE_NL:
		nlMonth := [...]string{"januari", "februari", "maart",
			"april", "mei", "juni", "juli", "augustus", "september",
			"oktober", "november", "december",
		}
		return fmt.Sprintf("%d %s %d", p.Date.Day(), nlMonth[p.Date.Month()-1], p.Date.Year())
	case format == DATE_LONG && p.Language == LANGUAGE_EN:
		return p.Date.Format("January 2, 2006")
	case format == DATE_FORMAL:
		return p.Date.Format(time.RFC1123Z)
	case format == DATE_SHORT:
		fallthrough
	default:
		return p.Date.Format("2006-01-02 00:00:00")
	}
}

func truncateOnSpace(text string, maxChars int) string {
	if len(text) <= maxChars {
		return text
	}

	shortText := ""
	ss := strings.Split(text, " ")
	for _, s := range ss {
		if len(shortText+" "+s) > maxChars {
			break
		}
		shortText += " " + s
	}

	return shortText
}
