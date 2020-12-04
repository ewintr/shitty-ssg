package site

import (
	"strings"
	"time"
)

const (
	TITLE_PREFIX           = "= "
	SUBTITLE_PREFIX        = "== "
	SUBSUBTITLE_PREFIX     = "=== "
	PARAGRAPH_SEPARATOR    = "\n\n"
	LINE_SEPARATOR         = "\n"
	CODE_PREFIX            = "----\n"
	CODE_SUFFIX            = "\n----"
	LISTITEM_PREFIX        = "* "
	PARAGRAPH_CONTINUATION = "\n+\n"
)

func NewPost(text string) Post {
	post := Post{
		Language: LANGUAGE_EN,
		Tags:     []Tag{},
	}

	// split up blocks
	var pars []string
	for _, s := range strings.Split(text, PARAGRAPH_SEPARATOR) {
		if s == "" {
			continue
		}

		pars = append(pars, s)
	}

	// actually, code blocks are allowed to have empty lines
	var blocks []string
	var inCode bool
	var currentBlock string
	for _, par := range pars {
		switch {
		case strings.HasPrefix(par, CODE_PREFIX) && strings.HasSuffix(par, CODE_SUFFIX):
			blocks = append(blocks, par)
		case !inCode && strings.HasPrefix(par, CODE_PREFIX):
			inCode = true
			currentBlock = par
		case inCode && !strings.HasSuffix(par, CODE_SUFFIX):
			currentBlock += PARAGRAPH_SEPARATOR + par
		case inCode && strings.HasSuffix(par, CODE_SUFFIX):
			currentBlock += PARAGRAPH_SEPARATOR + par
			blocks = append(blocks, currentBlock)
			inCode = false
			currentBlock = ""
		default:
			blocks = append(blocks, par)
		}
	}

	// interpret the blocks
	for i, p := range blocks {
		switch {
		case i == 0 && strings.HasPrefix(p, TITLE_PREFIX):
			ParseHeader(p, &post)
		case strings.HasPrefix(p, SUBTITLE_PREFIX):
			p = strings.TrimSpace(p)
			s := strings.Split(p, SUBTITLE_PREFIX)
			if len(s) == 1 || s[1] == "" {

				continue
			}
			post.Content = append(post.Content, SubTitle(s[1]))
		case strings.HasPrefix(p, SUBSUBTITLE_PREFIX):
			p = strings.TrimSpace(p)
			s := strings.Split(p, SUBSUBTITLE_PREFIX)
			if len(s) == 1 || s[1] == "" {

				continue
			}
			post.Content = append(post.Content, SubSubTitle(s[1]))
		case isCodeBlock(p):
			post.Content = append(post.Content, parseCodeBlock(p))
		case strings.HasPrefix(p, LISTITEM_PREFIX):
			p = strings.TrimSpace(p)
			var items []ListItem
			for i, ti := range strings.Split(p, LISTITEM_PREFIX) {
				if i > 0 {
					inline := ParseInline(strings.TrimSpace(ti))
					items = append(items, ListItem(inline))
				}
			}
			post.Content = append(post.Content, List(items))

		default:
			p = strings.TrimSpace(p)
			post.Content = append(post.Content, Paragraph(ParseInline(p)))
		}
	}

	return post
}

func isCodeBlock(par string) bool {
	return strings.HasPrefix(par, CODE_PREFIX) && strings.HasSuffix(par, CODE_SUFFIX)
}

func parseCodeBlock(par string) CodeBlock {
	ss := strings.Split(par, "\n")
	ss = ss[1 : len(ss)-1]
	content := strings.Join(ss, "\n")

	return CodeBlock(content)
}

func ParseHeader(text string, post *Post) {
	text = strings.TrimSpace(text)
	lines := strings.Split(text, LINE_SEPARATOR)
	for i, l := range lines {
		switch {
		case i == 0:
			s := strings.Split(l, TITLE_PREFIX)
			post.Title = s[1]
		case isDate(l):
			date, _ := time.Parse("2006-01-02", l)
			post.Date = date
		case strings.HasPrefix(l, ":kind:"):
			s := strings.Split(l, ":")
			post.Kind = NewKind(strings.TrimSpace(s[2]))
		case strings.HasPrefix(l, ":language:"):
			s := strings.Split(l, ":")
			post.Language = NewLanguage(strings.TrimSpace(s[2]))
		case strings.HasPrefix(l, ":tags:"):
			s := strings.Split(l, ":")
			t := strings.Split(s[2], ",")
			for _, tag := range t {
				post.Tags = append(post.Tags, Tag(strings.TrimSpace(tag)))
			}
		default:
			post.Author = l
		}
	}
}

func isDate(text string) bool {
	if _, err := time.Parse("2006-01-02", text); err == nil {
		return true
	}

	return false
}

func ParseInline(text string) []InlineElement {
	var e []InlineElement

	ss := strings.Split(text, "")
	var buffer, curWord, prevChar string
	var strong, emp, code, linkTitle bool
	wordStart := true
	for _, s := range ss {
		switch {
		case (s == "_" && wordStart) || (s == "_" && emp):
			e = addElement(e, buffer+curWord, strong, emp, code)
			emp = !emp
			buffer = ""
			curWord = ""
		case s == "*":
			e = addElement(e, buffer+curWord, strong, emp, code)
			buffer = ""
			curWord = ""
			strong = !strong
		case s == "`":
			e = addElement(e, buffer+curWord, strong, emp, code)
			code = !code
			buffer = ""
			curWord = ""
		case s == "[" && prevChar != "":
			e = addElement(e, buffer, strong, emp, code)
			linkTitle = true
			curWord += s
		case s == "]" && linkTitle:
			e = addLink(e, curWord)
			buffer = ""
			curWord = ""
			linkTitle = false
		case s == " " && !linkTitle:
			buffer += curWord + " "
			curWord = ""
		default:
			curWord += s
		}
		prevChar = s
		wordStart = false
		if prevChar == " " {
			wordStart = true
		}
	}
	if len(buffer+curWord) > 0 {
		e = addElement(e, buffer+curWord, false, false, false)
	}

	return e
}

func addLink(ies []InlineElement, linkText string) []InlineElement {
	ss := strings.Split(linkText, "[")
	if len(ss) < 2 {
		ss = append(ss, "ERROR")
	}

	return append(ies, Link{url: ss[0], title: ss[1]})
}

func addElement(ies []InlineElement, text string, strong, emp, code bool) []InlineElement {
	if len(text) == 0 {
		return ies
	}

	var ne InlineElement
	switch {
	case code:
		ne = CodeText(text)
	case strong && emp:
		ne = StrongEmpText(text)
	case strong && !emp:
		ne = StrongText(text)
	case !strong && emp:
		ne = EmpText(text)
	default:
		ne = PlainText(text)
	}

	return append(ies, ne)
}
