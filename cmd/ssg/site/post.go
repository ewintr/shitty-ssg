package site

import (
	"errors"
	"fmt"
	"html"
	"path"
	"strconv"
	"strings"
	"time"

	"code.ewintr.nl/go-kit/slugify"
	"code.ewintr.nl/shitty-ssg/pkg/adoc"
)

var (
	ErrInvalidPost = errors.New("invalid post")
)

var pluralKind = map[Kind]string{
	KIND_NOTE:    "notes",
	KIND_STORY:   "stories",
	KIND_ARTICLE: "articles",
}

type Post struct {
	doc        *adoc.ADoc
	baseURL    string
	prefixPath bool
	Date       time.Time
	Kind       Kind
	Language   Language
	Tags       []Tag
}

func NewPost(config *SiteConfig, doc *adoc.ADoc) *Post {
	var tags []Tag
	for _, t := range doc.Tags {
		tags = append(tags, Tag(t))
	}
	return &Post{
		doc:        doc,
		baseURL:    config.BaseURL,
		prefixPath: config.PathsWithKind,
		Date:       doc.Date,
		Kind:       config.MapKind(doc.Kind),
		Language:   Language(doc.Language),
		Tags:       tags,
	}
}

func (p Post) Slug() string {
	return slugify.Slugify(p.doc.Title)
}

func (p *Post) Year() string {
	return strconv.Itoa(p.Date.Year())
}

func (p *Post) Link() string {
	link := "/"
	if p.prefixPath {
		link = path.Join(link, pluralKind[p.Kind])
	}
	return fmt.Sprintf("%s/", path.Join(link, p.Year(), p.Slug()))
}

func (p *Post) FullLink() string {
	return fmt.Sprintf("%s%s", p.baseURL, p.Link())
}

func (p *Post) HTMLSummary() *HTMLSummary {
	summary := ""
	if len(p.doc.Content) > 0 {
		summary = fmt.Sprintf("<p>%s...</p>", TruncateOnSpace(p.doc.Content[0].Text(), 300))
	}

	return &HTMLSummary{
		Link:      p.Link(),
		Title:     p.doc.Title,
		Language:  p.Language,
		DateLong:  FormatDate(p.Date, p.Language, DATE_LONG),
		DateShort: FormatDate(p.Date, p.Language, DATE_SHORT),
		Summary:   summary,
	}
}

func (p *Post) HTMLPost() *HTMLPost {
	var content string
	for _, be := range p.doc.Content {
		content += fmt.Sprintf("%s\n", FormatBlock(be))
	}

	return &HTMLPost{
		Slug:      p.Slug(),
		Title:     html.EscapeString(p.doc.Title),
		DateLong:  FormatDate(p.Date, p.Language, DATE_LONG),
		DateShort: FormatDate(p.Date, p.Language, DATE_SHORT),
		Content:   content,
	}
}

func (p *Post) XMLPost() *XMLPost {
	var content string
	for _, be := range p.doc.Content {
		content += fmt.Sprintf("%s\n", FormatBlock(be))
	}

	return &XMLPost{
		Link:       p.FullLink(),
		Title:      html.EscapeString(p.doc.Title),
		DateFormal: FormatDate(p.Date, p.Language, DATE_FORMAL),
		Content:    content,
	}
}

func FormatDate(date time.Time, language Language, format DateFormat) string {
	switch {
	case format == DATE_LONG && language == LANGUAGE_NL:
		nlMonth := [...]string{"januari", "februari", "maart",
			"april", "mei", "juni", "juli", "augustus", "september",
			"oktober", "november", "december",
		}
		return fmt.Sprintf("%d %s %d", date.Day(), nlMonth[date.Month()-1], date.Year())
	case format == DATE_LONG && language == LANGUAGE_EN:
		return date.Format("January 2, 2006")
	case format == DATE_FORMAL:
		return date.Format(time.RFC1123Z)
	case format == DATE_SHORT:
		fallthrough
	default:
		return date.Format("2006-01-02 00:00:00")
	}
}

func TruncateOnSpace(text string, maxChars int) string {
	if len(text) <= maxChars {
		return text
	}

	var keep []string
	ss := strings.Split(text, " ")
	for _, s := range ss {
		if len(strings.Join(keep, " ")+s) > maxChars {
			break
		}
		keep = append(keep, s)
	}

	return strings.Join(keep, " ")
}
