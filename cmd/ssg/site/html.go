package site

import (
	"fmt"
	"html"
	"strings"

	"git.sr.ht/~ewintr/shitty-ssg/pkg/adoc"
)

type HTMLPost struct {
	Slug      string
	Title     string
	DateLong  string
	DateShort string
	Content   string
}

type HTMLSummary struct {
	Link      string
	Title     string
	Language  Language
	DateShort string
	DateLong  string
	Summary   string
}

func FormatBlock(block adoc.BlockElement) string {
	switch block.(type) {
	case adoc.Paragraph:
		text := ""
		for _, inline := range block.(adoc.Paragraph) {
			text += FormatInline(inline)
		}
		return fmt.Sprintf("<p>%s</p>", text)
	case adoc.SubTitle:
		return fmt.Sprintf("<h2>%s</h2>", html.EscapeString(block.Text()))
	case adoc.SubSubTitle:
		return fmt.Sprintf("<h3>%s</h3>", html.EscapeString(block.Text()))
	case adoc.CodeBlock:
		return fmt.Sprintf("<pre><code>%s</code></pre>", html.EscapeString(block.Text()))
	case adoc.List:
		var items []string
		for _, item := range block.(adoc.List) {
			itemText := ""
			for _, inline := range item {
				itemText += FormatInline(inline)
			}
			items = append(items, fmt.Sprintf("<li>%s</li>", itemText))
		}
		return fmt.Sprintf("<ul>\n%s\n</ul>", strings.Join(items, "\n"))
	default:
		return ""
	}
}

func FormatInline(src adoc.InlineElement) string {
	text := html.EscapeString(src.Text())
	switch src.(type) {
	case adoc.PlainText:
		return text
	case adoc.StrongText:
		return fmt.Sprintf("<strong>%s</strong>", text)
	case adoc.EmpText:
		return fmt.Sprintf("<em>%s</em>", text)
	case adoc.StrongEmpText:
		return fmt.Sprintf("<strong><em>%s</em></strong>", text)
	case adoc.Link:
		return fmt.Sprintf("<a href=%q>%s</a>", src.(adoc.Link).URL(), text)
	case adoc.CodeText:
		return fmt.Sprintf("<code>%s</code>", text)
	default:
		return ""
	}
}
