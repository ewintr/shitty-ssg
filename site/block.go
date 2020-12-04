package site

import (
	"fmt"
	"html"
	"strings"
)

type BlockElement interface {
	Text() string
	BlockHTML() string
}

type Paragraph []InlineElement

func (p Paragraph) Text() string {
	var text []string
	for _, ie := range p {
		text = append(text, ie.Text())
	}

	return strings.Join(text, " ")
}

func (p Paragraph) BlockHTML() string {
	var body string
	for _, ie := range p {
		body += ie.InlineHTML()
	}

	return fmt.Sprintf("<p>%s</p>", body)
}

type SubTitle string

func (st SubTitle) Text() string { return string(st) }
func (st SubTitle) BlockHTML() string {
	return fmt.Sprintf("<h2>%s</h2>", st)
}

type SubSubTitle string

func (st SubSubTitle) Text() string { return string(st) }
func (st SubSubTitle) BlockHTML() string {
	return fmt.Sprintf("<h3>%s</h3>", st)
}

type CodeBlock string

func (cb CodeBlock) Text() string { return string(cb) }
func (cb CodeBlock) BlockHTML() string {
	return fmt.Sprintf("<pre><code>%s</code></pre>", html.EscapeString(string(cb)))
}

type ListItem []InlineElement

func (li ListItem) Text() string {
	var text []string
	for _, ie := range li {
		text = append(text, ie.Text())
	}

	return fmt.Sprintf("%s%s", LISTITEM_PREFIX, strings.Join(text, " "))
}

func (li ListItem) HTML() string {
	var body string
	for _, ie := range li {
		body += ie.InlineHTML()
	}

	return fmt.Sprintf("<li>%s</li>", body)
}

type List []ListItem

func (l List) Text() string {
	var items []string
	for _, item := range l {
		items = append(items, item.Text())
	}

	return strings.Join(items, "\n")
}

func (l List) BlockHTML() string {
	var items []string
	for _, item := range l {
		items = append(items, item.HTML())
	}

	return fmt.Sprintf("<ul>\n%s\n</ul>", strings.Join(items, "\n"))
}
