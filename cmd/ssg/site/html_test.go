package site_test

import (
	"testing"

	"git.sr.ht/~ewintr/go-kit/test"
	"git.sr.ht/~ewintr/shitty-ssg/cmd/ssg/site"
	"git.sr.ht/~ewintr/shitty-ssg/pkg/adoc"
)

func TestFormatBlock(t *testing.T) {
	for _, tc := range []struct {
		name    string
		element adoc.BlockElement
		exp     string
	}{
		{
			name: "paragraph",
			element: adoc.Paragraph{
				adoc.PlainText("one"),
				adoc.PlainText("two"),
				adoc.PlainText("three"),
			},
			exp: "<p>onetwothree</p>",
		},
		{
			name:    "subtitle",
			element: adoc.SubTitle("text"),
			exp:     "<h2>text</h2>",
		},
		{
			name:    "subsubtitle",
			element: adoc.SubSubTitle("text"),
			exp:     "<h3>text</h3>",
		},
		{
			name:    "code",
			element: adoc.CodeBlock("text"),
			exp:     "<pre><code>text</code></pre>",
		},
		{
			name: "list",
			element: adoc.List{
				{adoc.PlainText("one")},
				{adoc.PlainText("two")},
				{adoc.PlainText("three")},
			},
			exp: "<ul>\n<li>one</li>\n<li>two</li>\n<li>three</li>\n</ul>",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			test.Equals(t, tc.exp, site.FormatBlock(tc.element))
		})
	}
}

func TestFormatInline(t *testing.T) {
	for _, tc := range []struct {
		name    string
		element adoc.InlineElement
		exp     string
	}{
		{
			name:    "plain text",
			element: adoc.PlainText("text"),
			exp:     "text",
		},
		{
			name:    "strong",
			element: adoc.StrongText("text"),
			exp:     "<strong>text</strong>",
		},
		{
			name:    "emphasis",
			element: adoc.EmpText("text"),
			exp:     "<em>text</em>",
		},
		{
			name:    "strong emphasis",
			element: adoc.StrongEmpText("text"),
			exp:     "<strong><em>text</em></strong>",
		},
		{
			name:    "link",
			element: adoc.NewLink("url", "title"),
			exp:     `<a href="url">title</a>`,
		},
		{
			name:    "code",
			element: adoc.CodeText("text"),
			exp:     "<code>text</code>",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			test.Equals(t, tc.exp, site.FormatInline(tc.element))
		})
	}
}
