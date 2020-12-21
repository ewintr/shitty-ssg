package site_test

import (
	"testing"
	"time"

	"git.sr.ht/~ewintr/go-kit/test"
	"git.sr.ht/~ewintr/shitty-ssg/cmd/ssg/site"
)

func TestPosts(t *testing.T) {
	kind1, kind2 := site.Kind("kind1"), site.Kind("kind2")
	tag1, tag2 := site.Tag("tag1"), site.Tag("tag2")
	post1 := &site.Post{
		Date: time.Date(2020, 12, 1, 0, 0, 0, 0, time.UTC),
		Kind: kind1,
		Tags: []site.Tag{tag1},
	}
	post2 := &site.Post{
		Date: time.Date(2019, 12, 1, 0, 0, 0, 0, time.UTC),
		Kind: kind2,
		Tags: []site.Tag{tag1, tag2},
	}
	post3 := &site.Post{
		Date: time.Date(2018, 12, 1, 0, 0, 0, 0, time.UTC),
		Kind: kind1,
		Tags: []site.Tag{tag2},
	}

	t.Run("sort", func(t *testing.T) {
		for _, tc := range []struct {
			name  string
			posts site.Posts
			exp   site.Posts
		}{
			{
				name:  "ordered",
				posts: site.Posts{post1, post2, post3},
			},
			{
				name:  "unordered",
				posts: site.Posts{post2, post3, post1},
			},
		} {
			t.Run(tc.name, func(t *testing.T) {
				exp := site.Posts{post1, post2, post3}
				test.Equals(t, exp, tc.posts.Sort())
			})
		}
	})

	posts := site.Posts{post1, post2, post3}

	t.Run("filter kind", func(t *testing.T) {
		test.Equals(t, site.Posts{post1, post3}, posts.FilterByKind(kind1))
	})

	t.Run("filter year", func(t *testing.T) {
		test.Equals(t, site.Posts{post2}, posts.FilterByYear("2019"))
	})

	t.Run("tilter tag", func(t *testing.T) {
		test.Equals(t, site.Posts{post2, post3}, posts.FilterByTag(tag2))
	})

	t.Run("limit", func(t *testing.T) {
		test.Equals(t, site.Posts{post1, post2}, posts.Limit(2))
	})

	t.Run("year list", func(t *testing.T) {
		test.Equals(t, []string{"2018", "2019", "2020"}, posts.YearList())
	})

	t.Run("tag list", func(t *testing.T) {
		test.Equals(t, []string{"tag1", "tag2"}, posts.TagList())
	})
}
