package site_test

import (
	"fmt"
	"testing"

	"git.sr.ht/~ewintr/go-kit/test"
	"git.sr.ht/~ewintr/shitty-ssg/site"
)

func TestPlainText(t *testing.T) {
	text := "text"
	pt := site.PlainText(text)

	test.Equals(t, text, pt.InlineHTML())
}

func TestStrongText(t *testing.T) {
	text := "text"
	st := site.StrongText(text)
	exp := fmt.Sprintf("<strong>%s</strong>", text)

	test.Equals(t, exp, st.InlineHTML())
}

func TestEmpText(t *testing.T) {
	text := "text"
	et := site.EmpText(text)
	exp := fmt.Sprintf("<em>%s</em>", text)

	test.Equals(t, exp, et.InlineHTML())
}

func TestLink(t *testing.T) {
	url := "http://example.com"
	title := "link title"
	link := site.NewLink(url, title)
	exp := fmt.Sprintf("<a href=%q>%s</a>", url, title)

	test.Equals(t, exp, link.InlineHTML())
}
