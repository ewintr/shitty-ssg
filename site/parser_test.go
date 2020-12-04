package site_test

import (
	"fmt"
	"testing"
	"time"

	"git.sr.ht/~ewintr/go-kit/test"
	"git.sr.ht/~ewintr/shitty-ssg/site"
)

func TestNewPost(t *testing.T) {
	one := "one"
	two := "two"
	three := "three"
	ptOne := site.PlainText(one)
	ptTwo := site.PlainText(two)
	ptThree := site.PlainText(three)
	for _, tc := range []struct {
		name  string
		input string
		exp   site.Post
	}{
		{
			name: "empty",
			exp: site.Post{
				Tags:     []site.Tag{},
				Language: site.LANGUAGE_EN,
			},
		},
		{
			name:  "title",
			input: "= Title",
			exp: site.Post{
				Title:    "Title",
				Tags:     []site.Tag{},
				Language: site.LANGUAGE_EN,
			},
		},
		{
			name:  "header",
			input: "= Title\nT. Test\n2020-10-27\n:tags:\ttag1, tag2\n:kind:\tnote\n:language:\tnl",
			exp: site.Post{
				Title:    "Title",
				Author:   "T. Test",
				Kind:     site.KIND_NOTE,
				Language: site.LANGUAGE_NL,
				Tags: []site.Tag{
					site.Tag("tag1"),
					site.Tag("tag2"),
				},
				Date: time.Date(2020, time.October, 27, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name:  "paragraphs",
			input: fmt.Sprintf("%s\n\n%s\n\n%s", one, two, three),
			exp: site.Post{
				Tags:     []site.Tag{},
				Language: site.LANGUAGE_EN,
				Content: []site.BlockElement{
					site.Paragraph([]site.InlineElement{ptOne}),
					site.Paragraph([]site.InlineElement{ptTwo}),
					site.Paragraph([]site.InlineElement{ptThree}),
				},
			},
		},
		{
			name:  "subtitle",
			input: "== Subtitle",
			exp: site.Post{
				Tags:     []site.Tag{},
				Language: site.LANGUAGE_EN,
				Content: []site.BlockElement{
					site.SubTitle("Subtitle"),
				},
			},
		},
		{
			name:  "code block",
			input: "----\nsome code\nmore code\n----",
			exp: site.Post{
				Tags:     []site.Tag{},
				Language: site.LANGUAGE_EN,
				Content: []site.BlockElement{
					site.CodeBlock("some code\nmore code"),
				},
			},
		},
		{
			name:  "code block with empty lines",
			input: "----\nsome code\n\nmore code\n----",
			exp: site.Post{
				Tags:     []site.Tag{},
				Language: site.LANGUAGE_EN,
				Content: []site.BlockElement{
					site.CodeBlock("some code\n\nmore code"),
				},
			},
		},
		{
			name:  "list",
			input: "* item 1\n* item 2\n* *item 3*\n",
			exp: site.Post{
				Tags:     []site.Tag{},
				Language: site.LANGUAGE_EN,
				Content: []site.BlockElement{
					site.List{
						site.ListItem([]site.InlineElement{site.PlainText("item 1")}),
						site.ListItem([]site.InlineElement{site.PlainText("item 2")}),
						site.ListItem([]site.InlineElement{site.StrongText("item 3")}),
					},
				},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			act := site.NewPost(tc.input)

			test.Equals(t, tc.exp, act)
		})
	}
}

func TestPostBodyHTML(t *testing.T) {
	text := `= Title

Some text. With some *strong*. And a http://example.com[link].

== A Sub Title

And some more text.
	`
	post := site.NewPost(text)
	act := post.HTMLPost()

	exp := &site.HTMLPost{
		Slug:      "title",
		Title:     "Title",
		DateLong:  "January 1, 0001",
		DateShort: "0001-01-01 00:00:00",
		Content: `<p>Some text. With some <strong>strong</strong>. And a <a href="http://example.com">link</a>.</p>
<h2>A Sub Title</h2>
<p>And some more text.</p>
`,
	}

	test.Equals(t, exp, act)

}

func TestParseInline(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		exp   []site.InlineElement
	}{{
		name: "empty",
	},
		{
			name:  "plain",
			input: "some test text",
			exp: []site.InlineElement{
				site.PlainText("some test text")},
		},
		{
			name:  "strong",
			input: "*some strong text*",
			exp: []site.InlineElement{
				site.StrongText("some strong text"),
			},
		},
		{
			name:  "strong in plain",
			input: "some *strong* text",
			exp: []site.InlineElement{
				site.PlainText("some "),
				site.StrongText("strong"),
				site.PlainText(" text"),
			},
		},
		{
			name:  "emphasis",
			input: "_some emphasized text_",
			exp: []site.InlineElement{
				site.EmpText("some emphasized text"),
			},
		},
		{
			name:  "emphasis in plain",
			input: "some _emphasized_ text",
			exp: []site.InlineElement{
				site.PlainText("some "),
				site.EmpText("emphasized"),
				site.PlainText(" text"),
			},
		},
		{
			name:  "emp and strong in plain",
			input: "some _*special*_ text",
			exp: []site.InlineElement{
				site.PlainText("some "),
				site.StrongEmpText("special"),
				site.PlainText(" text"),
			},
		},
		{
			name:  "link",
			input: "a link[title] somewhere",
			exp: []site.InlineElement{
				site.PlainText("a "),
				site.NewLink("link", "title"),
				site.PlainText(" somewhere"),
			},
		},
		{
			name:  "code",
			input: "`command`",
			exp: []site.InlineElement{
				site.CodeText("command"),
			},
		},
		{
			name:  "code in plain",
			input: "some `code` in text",
			exp: []site.InlineElement{
				site.PlainText("some "),
				site.CodeText("code"),
				site.PlainText(" in text"),
			},
		},
		{
			name:  "link with underscore",
			input: "https://example.com/some_url[some url]",
			exp: []site.InlineElement{
				site.NewLink("https://example.com/some_url", "some url"),
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			act := site.ParseInline(tc.input)

			test.Equals(t, tc.exp, act)

		})
	}
}
