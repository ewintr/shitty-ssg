package site_test

import (
	"testing"
	"time"

	"go-mod.ewintr.nl/go-kit/test"
	"go-mod.ewintr.nl/shitty-ssg/cmd/ssg/site"
	"go-mod.ewintr.nl/shitty-ssg/pkg/adoc"
)

func TestPost(t *testing.T) {
	docKind := adoc.KIND_NOTE
	siteKind := site.Kind(docKind)
	config := &site.SiteConfig{
		BaseURL: "base_url",
		KindMap: map[adoc.Kind]site.Kind{
			docKind: siteKind,
		},
		PathsWithKind: true,
	}
	title := "title thing"
	author := "author"
	language := adoc.LANGUAGE_EN
	path := "/path"
	date := time.Date(2020, 12, 28, 7, 23, 45, 0, time.UTC)
	tag1, tag2 := adoc.Tag("tag1"), adoc.Tag("tag2")
	par1 := adoc.Paragraph{adoc.PlainText("one")}
	par2 := adoc.Paragraph{adoc.PlainText("two")}
	post := site.NewPost(config, &adoc.ADoc{
		Title:    title,
		Author:   author,
		Kind:     docKind,
		Language: language,
		Path:     path,
		Date:     date,
		Tags:     []adoc.Tag{tag1, tag2},
		Content:  []adoc.BlockElement{par1, par2},
	})

	t.Run("new", func(t *testing.T) {
		test.Equals(t, date, post.Date)
		test.Equals(t, siteKind, post.Kind)
		test.Equals(t, site.Language(language), post.Language)
		test.Equals(t, []site.Tag{site.Tag(tag1), site.Tag(tag2)}, post.Tags)
	})

	t.Run("tags", func(t *testing.T) {
		test.Equals(t, []site.Tag{site.Tag("tag1"), site.Tag("tag2")}, post.Tags)
	})

	t.Run("slug", func(t *testing.T) {
		test.Equals(t, "title-thing", post.Slug())
	})

	t.Run("year", func(t *testing.T) {
		test.Equals(t, "2020", post.Year())
	})

	t.Run("link", func(t *testing.T) {
		test.Equals(t, "/notes/2020/title-thing/", post.Link())
	})

	t.Run("full link", func(t *testing.T) {
		test.Equals(t, "base_url/notes/2020/title-thing/", post.FullLink())
	})

	t.Run("html summary", func(t *testing.T) {
		exp := &site.HTMLSummary{
			Link:      "/notes/2020/title-thing/",
			Title:     "title thing",
			Language:  site.LANGUAGE_EN,
			DateShort: "2020-12-28 00:00:00",
			DateLong:  "December 28, 2020",
			Summary:   "<p>one...</p>",
		}

		test.Equals(t, exp, post.HTMLSummary())
	})

	t.Run("html post", func(t *testing.T) {
		exp := &site.HTMLPost{
			Slug:      "title-thing",
			Title:     "title thing",
			DateLong:  "December 28, 2020",
			DateShort: "2020-12-28 00:00:00",
			Content:   "<p>one</p>\n<p>two</p>\n",
		}

		test.Equals(t, exp, post.HTMLPost())
	})

	t.Run("xml post", func(t *testing.T) {
		exp := &site.XMLPost{
			Link:       "base_url/notes/2020/title-thing/",
			Title:      "title thing",
			DateFormal: "Mon, 28 Dec 2020 07:23:45 +0000",
			Content:    "<p>one</p>\n<p>two</p>\n",
		}

		test.Equals(t, exp, post.XMLPost())
	})
}

func TestFormatDate(t *testing.T) {
	date := time.Date(2020, 12, 28, 7, 23, 45, 0, time.UTC)
	for _, tc := range []struct {
		name     string
		language site.Language
		format   site.DateFormat
		exp      string
	}{
		{
			name:     "long nl",
			language: site.LANGUAGE_NL,
			format:   site.DATE_LONG,
			exp:      "28 december 2020",
		},
		{
			name:     "long en",
			language: site.LANGUAGE_EN,
			format:   site.DATE_LONG,
			exp:      "December 28, 2020",
		},
		{
			name:     "formal nl",
			language: site.LANGUAGE_NL,
			format:   site.DATE_FORMAL,
			exp:      "Mon, 28 Dec 2020 07:23:45 +0000",
		},
		{
			name:     "formal en",
			language: site.LANGUAGE_EN,
			format:   site.DATE_FORMAL,
			exp:      "Mon, 28 Dec 2020 07:23:45 +0000",
		},
		{
			name:     "short nl",
			language: site.LANGUAGE_NL,
			format:   site.DATE_SHORT,
			exp:      "2020-12-28 00:00:00",
		},
		{
			name:     "short en",
			language: site.LANGUAGE_EN,
			format:   site.DATE_SHORT,
			exp:      "2020-12-28 00:00:00",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			test.Equals(t, tc.exp, site.FormatDate(date, tc.language, tc.format))
		})
	}
}

func TestTruncateOnSpace(t *testing.T) {
	text := "this is a short text"
	for _, tc := range []struct {
		name string
		text string
		max  int
		exp  string
	}{
		{
			name: "empty",
		},
		{
			name: "short text",
			text: text,
			max:  150,
			exp:  text,
		},
		{
			name: "truncate on space",
			text: text,
			max:  10,
			exp:  "this is a",
		},
		{
			name: "truncate in word",
			text: text,
			max:  12,
			exp:  "this is a",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			test.Equals(t, tc.exp, site.TruncateOnSpace(tc.text, tc.max))
		})
	}
}
