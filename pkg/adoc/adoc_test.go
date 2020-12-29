package adoc_test

import (
	"testing"

	"git.sr.ht/~ewintr/go-kit/test"
	"git.sr.ht/~ewintr/shitty-ssg/pkg/adoc"
)

func TestNewKind(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		exp   adoc.Kind
	}{
		{
			name: "empty",
			exp:  adoc.KIND_UNKNOWN,
		},
		{
			name:  "unknown",
			input: "something",
			exp:   adoc.KIND_UNKNOWN,
		},
		{
			name:  "note",
			input: "note",
			exp:   adoc.KIND_NOTE,
		},
		{
			name:  "vkv",
			input: "vkv",
			exp:   adoc.KIND_VKV,
		},
		{
			name:  "story",
			input: "verhaal",
			exp:   adoc.KIND_STORY,
		},
		{
			name:  "snippet",
			input: "los",
			exp:   adoc.KIND_SNIPPET,
		},
		{
			name:  "essay",
			input: "essay",
			exp:   adoc.KIND_ESSAY,
		},
		{
			name:  "tutorial",
			input: "tutorial",
			exp:   adoc.KIND_TUTORIAL,
		},
		{
			name:  "article",
			input: "article",
			exp:   adoc.KIND_ARTICLE,
		},
		{
			name:  "work note",
			input: "work",
			exp:   adoc.KIND_WORK,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			act := adoc.NewKind(tc.input)
			test.Equals(t, tc.exp, act)
		})
	}
}

func TestNewLanguage(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		exp   adoc.Language
	}{
		{
			name: "empty",
			exp:  adoc.LANGUAGE_UNKNOWN,
		},
		{
			name:  "dutch lower",
			input: "nl",
			exp:   adoc.LANGUAGE_NL,
		},
		{
			name:  "dutch upper",
			input: "NL",
			exp:   adoc.LANGUAGE_NL,
		},
		{
			name:  "english lower",
			input: "en",
			exp:   adoc.LANGUAGE_EN,
		},
		{
			name:  "english upper",
			input: "EN",
			exp:   adoc.LANGUAGE_EN,
		},
		{
			name:  "unknown",
			input: "something",
			exp:   adoc.LANGUAGE_UNKNOWN,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			act := adoc.NewLanguage(tc.input)
			test.Equals(t, tc.exp, act)
		})

	}
}
