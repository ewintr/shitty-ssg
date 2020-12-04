package site

import (
	"fmt"
	"html"
)

type InlineType int

type InlineElement interface {
	Text() string
	InlineHTML() string
}

type PlainText string

func (pt PlainText) Text() string { return string(pt) }

func (pt PlainText) InlineHTML() string {
	return html.EscapeString(string(pt))
}

type StrongText string

func (st StrongText) Text() string { return string(st) }

func (st StrongText) InlineHTML() string {
	return fmt.Sprintf("<strong>%s</strong>", html.EscapeString(string(st)))
}

type EmpText string

func (et EmpText) Text() string { return string(et) }

func (et EmpText) InlineHTML() string {
	return fmt.Sprintf("<em>%s</em>", html.EscapeString(string(et)))
}

type StrongEmpText string

func (set StrongEmpText) Text() string { return string(set) }

func (set StrongEmpText) InlineHTML() string {
	return fmt.Sprintf("<strong><em>%s</em></strong>", html.EscapeString(string(set)))
}

type Link struct {
	url   string
	title string
}

func NewLink(url, title string) Link {
	return Link{
		url:   url,
		title: title,
	}
}

func (l Link) Text() string { return l.title }

func (l Link) InlineHTML() string {
	return fmt.Sprintf("<a href=%q>%s</a>", l.url, html.EscapeString(l.title))
}

type CodeText string

func (ct CodeText) Text() string { return string(ct) }

func (ct CodeText) InlineHTML() string {
	return fmt.Sprintf("<code>%s</code>", html.EscapeString(string(ct)))
}
