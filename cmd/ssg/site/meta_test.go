package site_test

import (
	"testing"

	"go-mod.ewintr.nl/go-kit/test"
	"go-mod.ewintr.nl/shitty-ssg/cmd/ssg/site"
	"go-mod.ewintr.nl/shitty-ssg/pkg/adoc"
)

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
