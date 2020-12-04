package site

import (
	"fmt"
	"testing"

	"git.sr.ht/~ewintr/go-kit/test"
)

func TestParagraph(t *testing.T) {
	p := Paragraph([]InlineElement{
		PlainText("one "),
		PlainText("two "),
		PlainText("three"),
	})

	exp := "<p>one two three</p>"
	test.Equals(t, exp, p.BlockHTML())
}

func TestSubTitle(t *testing.T) {
	text := "text"
	st := SubTitle(text)

	exp := fmt.Sprintf("<h2>%s</h2>", text)
	test.Equals(t, exp, st.BlockHTML())
}
