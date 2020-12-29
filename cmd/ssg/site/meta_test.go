package site_test

import (
	"testing"

	"git.sr.ht/~ewintr/go-kit/test"
	"git.sr.ht/~ewintr/shitty-ssg/cmd/ssg/site"
	"git.sr.ht/~ewintr/shitty-ssg/pkg/adoc"
)

func TestNewKind(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input adoc.Kind
		exp   site.Kind
	}{
		{
			name: "empty",
			exp:  site.KIND_INVALID,
		},
		{
			name:  "note",
			input: adoc.KIND_NOTE,
			exp:   site.KIND_NOTE,
		},
		{
			name:  "work",
			input: adoc.KIND_WORK,
			exp:   site.KIND_INVALID,
		},
		{
			name:  "vkv",
			input: adoc.KIND_VKV,
			exp:   site.KIND_STORY,
		},
		{
			name:  "essay",
			input: adoc.KIND_ESSAY,
			exp:   site.KIND_ARTICLE,
		},
		{
			name:  "article",
			input: adoc.KIND_ARTICLE,
			exp:   site.KIND_ARTICLE,
		},
		{
			name:  "tutorial",
			input: adoc.KIND_TUTORIAL,
			exp:   site.KIND_ARTICLE,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			test.Equals(t, tc.exp, site.NewKind(tc.input))
		})
	}
}

func TestNewLanguage(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input adoc.Language
		exp   site.Language
	}{
		{
			name:  "nl",
			input: adoc.LANGUAGE_NL,
			exp:   site.LANGUAGE_NL,
		},
		{
			name:  "en",
			input: adoc.LANGUAGE_EN,
			exp:   site.LANGUAGE_EN,
		},
		{
			name:  "unknown",
			input: adoc.LANGUAGE_UNKNOWN,
			exp:   site.LANGUAGE_EN,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			test.Equals(t, tc.exp, site.NewLanguage(tc.input))
		})
	}
}
